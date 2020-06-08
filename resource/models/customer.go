package models

//Customers ...
type Customers struct {
	Base
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}
