package models

//Categories ...
type Categories struct {
	Base
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desc"`
}
