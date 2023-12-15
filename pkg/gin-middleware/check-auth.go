package middleware

import (
	"github.com/gin-gonic/gin"
	jwtLib "github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

type AuthInterface interface {
	Validate(tokenString string) (claims jwtLib.MapClaims, err error)
}

// CheckAuth - a function that checks the authentication method and validates the token.
// It takes the request object and the JWT instance as parameters.
// It returns an error if the authentication method is unsupported, the token is invalid, or the token is missing.
func CheckAuth(authService AuthInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header from the request
		authHeader := c.GetHeader("Authorization")

		// Check if the Authorization header is present
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}

		// Split the Authorization header into the authentication scheme and the token
		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}

		// Check the authentication scheme
		authScheme := authParts[0]
		token := authParts[1]

		if authScheme == "Bearer" {
			// Validate the token using the Validate function from the jwt package
			_, err := authService.Validate(token)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unsupported authentication scheme"})
			return
		}

		// Continue to the next middleware or handler
		c.Next()
	}
}
