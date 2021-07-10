package repository

import (
	"errors"

	"github.com/mhope-2/currency_converter/database/models"
)

// get source currency
func (r *Repository) GetSourceCurrency(source string) (*models.Currency, error){

	currency := models.Currency{}

	if source := r.DB.First(&currency, "code = ?", source).Error; source != nil {
		return nil, source
	}

	if currency.ID == 0 {
	    return nil, errors.New("Currency Doesn't Exist")
	}

	return &currency, nil
}


// get source currency
func (r *Repository) GetTargetCurrency(target string) (*models.Currency, error){

	currency := models.Currency{}

	if source := r.DB.First(&currency, "code = ?", target).Error; source != nil {
		return nil, source
	}

	if currency.ID == 0 {
	    return nil, errors.New("Currency Doesn't Exist")
	}

	return &currency, nil
}


// get source currency
func (r *Repository) GetExchangeRate(sourceID uint, targetID uint) (*models.ExchangeRates, error){

	exchangeRate := models.ExchangeRates{}

	if exchangeRateInstance := r.DB.First(&exchangeRate, "source_currency_id = ? AND target_currency_id = ?", 
										  sourceID, targetID).Error; exchangeRateInstance != nil {
		return nil, exchangeRateInstance
	}

	return &exchangeRate, nil
}


// get source currency
func (r *Repository) GetConvertedAmount(exchangeRate float64, requestAmount float64) (float64, error){

	convertedAmount := exchangeRate * float64(requestAmount)

	return convertedAmount,nil
}
