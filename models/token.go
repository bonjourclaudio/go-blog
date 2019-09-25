package models

import "github.com/dgrijalva/jwt-go"

//Token struct declaration
type Token struct {
	UserID uint
	Firstname   string
	Lastname string
	Email  string
	*jwt.StandardClaims
}