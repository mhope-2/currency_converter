package models

type ExchangeRates struct {
	Default
	Currency Currency
	SourceID Currency `json:"source_id"`
	TargetID Currency `json:"target_id"`
	Value 	 float64  `json:"value"`
	Inverse  float64  `json:"invertse"`
	Name     string   `json:"name"`
	Code     string   `json:"code"`
	Sign     string   `json:"sign"`
}