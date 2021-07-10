package handler

import (
	// "fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhope-2/currency_converter/database/models"
	"gorm.io/gorm"
)


type ConvertCurrencyRequest struct {
	Source    string `json:"source_currency" validate:"required"`
	Target    string `json:"target_currency" validate:"required"`
	Amount    float64 `json:"amount" validate:"required"`

}

// List all currencies
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

// List all exchange rates
func (h *Handler) ListExchangeRates(c *gin.Context) {
	var exchangeRates []models.ExchangeRates

	if err := h.DB.Find(&exchangeRates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to retrieve exchange rates",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response":   "success",
		"exchange_rates": exchangeRates,
	})
}


// convert currency
func (h *Handler) ConvertCurrency(c *gin.Context) {
	var request ConvertCurrencyRequest

	// currency := models.Currency{}

	if err := c.BindJSON(&request); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{

			"message": "failed to parse request",
		})
		return
	}

	sourceCurrency, err := h.Repo.GetSourceCurrency(request.Source)
	targetCurrency, err := h.Repo.GetTargetCurrency(request.Target)

	exchangeRateInstance, err := h.Repo.GetExchangeRate(sourceCurrency.ID, targetCurrency.ID)

	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Request failed",
		})
		return
	}

	if sourceCurrency != nil && targetCurrency != nil && exchangeRateInstance != nil{

		convertedAmout, err := h.Repo.GetConvertedAmount(exchangeRateInstance.Value, request.Amount)

		if err != nil{
			c.JSON(http.StatusOK, gin.H{
				"error": "Invalid Amount Passed",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"source_currency":  request.Source,
			"target_currency": request.Target,
			"amount": request.Amount,
			"converted_value": convertedAmout,
		})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "Either of the Currencies Passed Doesn't Exist",
		})
	}
}
