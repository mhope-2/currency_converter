package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"os"

	"github.com/gin-gonic/gin"
)


func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Run the other tests
	os.Exit(m.Run())
  }


func TestListCurrencies(t *testing.T) {
	
	// Create mock request
	req, err := http.NewRequest(http.MethodGet, "/v1/currencies", nil)

	if err != nil {
		t.Fatalf("Could not create request. Cause: %v\n", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Setup router
	router := gin.Default()

	// var h *Handler
	router.GET("/v1/currencies")

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("ERROR, expecting %v, got %v", http.StatusOK, w.Code)
	}

}
