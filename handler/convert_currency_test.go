package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)


func TestConvertCurrency(t *testing.T) {

	data := url.Values{}
    data.Set("source_currency", "NGN")
    data.Set("target_currency", "GHS")
	data.Set("amount", "100000")

	
	// Create mock request
	req, err := http.NewRequest("POST", "/convert/currency", strings.NewReader(data.Encode()))

	if err != nil {
		t.Fatalf("Could not create request. Cause: %v\n", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Setup router
	router := gin.Default()

	// var h *Handler
	router.POST("/convert/currency")

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("ERROR, expecting %v, got %v", http.StatusOK, w.Code)
	}

}
