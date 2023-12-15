package jwt

import (
	"fmt"
	jwtLib "github.com/golang-jwt/jwt/v5"
	"time"
)

type Jwt struct {
	exp    time.Duration
	secret string
}

// NewJwt - create new jwt struct.
// exp is the validity time.
// secret - secret string.
func NewJwt(exp time.Duration, secret string) *Jwt {
	return &Jwt{exp: exp, secret: secret}
}

// Generate - generate new JWT
func (o *Jwt) Generate(userId interface{}) (tokenString string, err error) {

	token := jwtLib.NewWithClaims(jwtLib.SigningMethodHS256, jwtLib.MapClaims{
		"exp":     time.Now().Add(o.exp).Unix(),
		"user_id": userId,
	})
	tokenString, err = token.SignedString([]byte(o.secret))

	return
}

// Validate - checks the validity of the token and returns its claims in case of success.
// If the token is invalid, an error is returned.
func (o *Jwt) Validate(tokenString string) (claims jwtLib.MapClaims, err error) {
	// Parse the token and verify its signature
	token, err := jwtLib.Parse(tokenString, func(token *jwtLib.Token) (interface{}, error) {
		// Check that the signing method is what we expect
		if _, ok := token.Method.(*jwtLib.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key
		return []byte(o.secret), nil
	})

	// Check for errors
	if err != nil {
		return nil, err
	}

	// Check that the token is valid
	if claims, ok := token.Claims.(jwtLib.MapClaims); ok && token.Valid {
		// Check if the token has expired
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return nil, fmt.Errorf("token has expired")
			}
		}

		fmt.Println(claims)
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
