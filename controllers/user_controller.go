package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetUserById(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{"route": "user route"})
}
