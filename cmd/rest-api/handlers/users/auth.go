package users

import (
	"errors"
	middleware "github.com/evzubkov/go-boilerplate/pkg/gin-middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type AuthRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type AuthResponse struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

// Auth - auth for user
// @Router /users/auth [get]
//
// @Summary auth for user
// @Description auth for user
// @Description
// @ID Auth
// @Tags Users
//
// @Param data body AuthRequest true "AuthRequest"
//
// @Success 200 {object} AuthRequest "Ok"
// @Failure 500 {object} AuthResponse "Error Msg"
func (o *Handler) Auth(ctx *gin.Context) {

	request := &AuthRequest{}
	if err := ctx.ShouldBind(request); err != nil && errors.As(err, &validator.ValidationErrors{}) {
		middleware.RenderBindingErrors(ctx, err.(validator.ValidationErrors))
		return
	}
}
