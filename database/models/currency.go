package models

type Currency struct {
	Default
	Name string `json:"name"`
	Code string `json:"code"`
	Sign string `json:"sign"`
}