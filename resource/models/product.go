package models

//Products ...
type Products struct {
	Base
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
