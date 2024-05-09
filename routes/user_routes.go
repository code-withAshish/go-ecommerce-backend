package routes

import (
	"backend/controllers"
	"github.com/labstack/echo/v4"
)

func UserSubRoutes(router *echo.Group) {
	router.GET("/getUserById", controllers.GetUserById)
}
