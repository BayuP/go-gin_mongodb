package models

//Products ...
type Products struct {
	Base
	ID         string `json:"id"`
	Name       string `json:"name"`
	CategoryID string `json:"category_id"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
}
