package models

import jwt "github.com/dgrijalva/jwt-go"

// Token ...
type Token struct {
	//UserID uint
	Username string
	//Email  string
	*jwt.StandardClaims
}
