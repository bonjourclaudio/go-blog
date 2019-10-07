package models

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	UserID uint
	Firstname   string
	Lastname string
	Email  string
	*jwt.StandardClaims
}