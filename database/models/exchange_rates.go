package models

type ExchangeRates struct {
	Model
	SourceCurrencyID   	   uint       `json:"source"`
	TargetCurrencyID       uint       `json:"target"`
	Value 	 			   float64    `json:"value"`
}

