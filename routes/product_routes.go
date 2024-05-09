package routes

import (
	"backend/controllers"
	"github.com/labstack/echo/v4"
)

func ProductSubRoutes(router *echo.Group) {
	router.GET("/getProductById", controllers.GetProductById)
}
