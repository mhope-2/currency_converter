package models

type ExchangeRates struct {
	Model
	Currency Currency
	Source Currency `json:"source_id"`
	Target Currency `json:"target_id"`
	Value 	 float64  `json:"value"`
	Inverse  float64  `json:"invertse"`
	Name     string   `json:"name"`
	Code     string   `json:"code"`
	Sign     string   `json:"sign"`
}