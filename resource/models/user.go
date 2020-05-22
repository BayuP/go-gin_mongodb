package models

//User models
type User struct {
	Base
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
