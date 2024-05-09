package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetProductById(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{"route": "product route controller"})
}
