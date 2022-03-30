package utils

import "github.com/dgrijalva/jwt-go"

type AccessToken struct {
	Token  string
	Claims jwt.StandardClaims
}
