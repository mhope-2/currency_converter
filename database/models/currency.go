package models

// currency model struct
type Currency struct {
	Model
	Name 	string `json:"name"`
	Code 	string `json:"code"`
	Symbol 	string `json:"symbol"`
}