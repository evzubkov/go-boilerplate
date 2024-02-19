package jwt_test

import (
	"github.com/evzubkov/go-boilerplate/pkg/jwt"
	"testing"
	"time"
)

func TestGenerateAndValidateToken(t *testing.T) {
	// Create a new JWT instance
	jwtInst := jwt.NewJwt(time.Minute, "secret")

	// Generate a token
	userId := "123"
	tokenString, err := jwtInst.Generate(userId)
	if err != nil {
		t.Errorf("Failed to generate token: %v", err)
	}

	// Validate the token
	claims, err := jwtInst.Validate(tokenString)
	if err != nil {
		t.Errorf("Failed to validate token: %v", err)
	}

	// Check the user_id claim
	if userIdTest, ok := claims["user_id"]; !ok || userIdTest != userId {
		t.Errorf("Invalid user_id claim: %v", userIdTest)
	}
}

func TestInvalidToken(t *testing.T) {
	// Create a new JWT instance
	jwtInst := jwt.NewJwt(time.Minute, "secret")

	// Validate an invalid token
	invalidToken := "invalid_token"
	_, err := jwtInst.Validate(invalidToken)
	if err == nil {
		t.Errorf("Expected error for invalid token, but got nil")
	}
}

func TestExpiredToken(t *testing.T) {
	// Create a new JWT instance with a short expiration time
	jwtInst := jwt.NewJwt(time.Second, "secret")

	// Generate a token
	userId := "123"
	tokenString, err := jwtInst.Generate(userId)
	if err != nil {
		t.Errorf("Failed to generate token: %v", err)
	}

	// Wait for the token to expire
	time.Sleep(2 * time.Second)

	// Validate the expired token
	_, err = jwtInst.Validate(tokenString)
	if err == nil {
		t.Errorf("Expected error for expired token, but got nil")
	}
}
