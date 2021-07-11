package models

// exchange rate model struct
type ExchangeRates struct {
	Model
	SourceCurrencyID   	   uint       `json:"source"`
	TargetCurrencyID       uint       `json:"target"`
	Value 	 	           float64    `json:"value"`
}

