package hello

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Hello Привет
// @Router /hello [get]
//
// @Summary Привет
// @Description Привет
// @Description
// @Description Ошибки:
// @Description "Internal Server Error" - что-то совсем пошло не так
// @ID Hello
// @Tags Hello
// @Accept json
// @Produce json
//
// @Success 200 {string} string "OK"
// @Failure 500 {object} ErrAnswer "Internal Server Error"
func (h *Handler) Hello(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]string{
		"msg": "Hello!",
	})
}
