package handler

import (

	"github.com/gin-gonic/gin"
	"github.com/mhope-2/currency_converter/repository"
	"gorm.io/gorm"

	)

type Handler struct {
	DB                   *gorm.DB
	Repo                 *repository.Repository
}


func New(DB *gorm.DB) *Handler {
	repo := repository.New(DB)

	return &Handler{
		DB:                   DB,
		Repo:                 repo,
	}
}

func (h *Handler) Register(v1 *gin.RouterGroup){

	currency := v1.Group("/currencies")
	currency.GET("/", h.ListCurrencies)

	exchange_rates := v1.Group("/exchange/rates")
	exchange_rates.GET("/", h.ListExchangeRates)

	convert_currency := v1.Group("/convert/currency")
	convert_currency.POST("/", h.ConvertCurrency)
}





