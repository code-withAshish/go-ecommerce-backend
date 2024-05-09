package controllers

import (
	"backend/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

var db = utils.GetDatabaseConnection()

func GetProductById(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{"route": "product route controller"})
}
