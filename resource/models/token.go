package models

import jwt "github.com/dgrijalva/jwt-go"

// Token ...
type Token struct {
	//UserID uint
	ID       string
	Username string
	RoleID   string
	//Email  string
	*jwt.StandardClaims
}
