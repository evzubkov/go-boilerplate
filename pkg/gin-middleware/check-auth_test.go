package middleware_test

import (
	"errors"
	"github.com/gin-gonic/gin"
	jwtLib "github.com/golang-jwt/jwt/v5"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"go-boilerplate/pkg/gin-middleware"
)

type mockAuth struct{}

func (m *mockAuth) Validate(tokenString string) (claims jwtLib.MapClaims, err error) {
	if tokenString == "valid-token" {
		return jwtLib.MapClaims{"user_id": "123"}, nil
	}
	return nil, errors.New("invalid token")
}

func TestCheckAuth_ValidToken(t *testing.T) {
	// Create a new mock AuthInterface instance
	auth := &mockAuth{}

	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "Bearer valid-token")
	c.Request = req

	// Call the CheckAuth function
	middleware.CheckAuth(auth)(c)

	// Assert that the response status code is 200
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckAuth_InvalidToken(t *testing.T) {
	// Create a new mock AuthInterface instance
	auth := &mockAuth{}

	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "Bearer invalid-token")
	c.Request = req

	// Call the CheckAuth function
	middleware.CheckAuth(auth)(c)

	// Assert that the response status code is 401
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestCheckAuth_MissingAuthorizationHeader(t *testing.T) {
	// Create a new mock AuthInterface instance
	auth := &mockAuth{}

	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	c.Request = req

	// Call the CheckAuth function
	middleware.CheckAuth(auth)(c)

	// Assert that the response status code is 401
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestCheckAuth_InvalidAuthorizationHeaderFormat(t *testing.T) {
	// Create a new mock AuthInterface instance
	auth := &mockAuth{}

	// Create a new gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "invalid-format")
	c.Request = req

	// Call the CheckAuth function
	middleware.CheckAuth(auth)(c)

	// Assert that the response status code is 401
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
