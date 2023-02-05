package utils

import "github.com/dgrijalva/jwt-go"

// jwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}
