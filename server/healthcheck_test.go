package server

import (
	"bytes"
	"os"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)
  
	// Run the other tests
	os.Exit(m.Run())
  }

  
func TestHealthCheck(t *testing.T){

	// Create mock request
	req, err := http.NewRequest(http.MethodGet, "/", nil)

	if err != nil{
		t.Fatalf("Could not create request. Cause: %v\n", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Setup router
	router := gin.Default()

	router.GET("/", HealthCheck)

	// Create a response recorder so you can inspect the response
    w := httptest.NewRecorder()

    // Perform the request
    router.ServeHTTP(w, req)
    
	response := bytes.NewBuffer([]byte("{\"status\":\"OK\"}"))

	res := bytes.Compare(w.Body.Bytes(), response.Bytes())

	if res != 0 && w.Code != http.StatusOK {
		t.Fatalf("ERROR, expecting %v, got %v", response, w.Body)
	}

}