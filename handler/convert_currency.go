package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhope-2/currency_converter.git/database/models"
	"github.com/mhope-2/go_book_api/database/models"
	"gorm.io/gorm"
)


type ConvertCurrencyRequest struct {
	Source    string `json:"source_currency" validate:"required"`
	Target string `json:"target_currency" validate:"required"`

}


func (h *Handler) ListCurrencies(c *gin.Context) {
	var currencies []models.Currency

	if err := h.DB.Find(&currencies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to retrieve currencies",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response":   "success",
		"currencies": currencies,
	})
}


func (h *Handler) ListExchangeRates(c *gin.Context) {
	var currencies []models.Currency

	if err := h.DB.Find(&currencies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to retrieve currencies",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response":   "success",
		"currencies": currencies,
	})
}

