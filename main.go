package main

import (
	"backend/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	routes.UserSubRoutes(app.Group("/user"))
	routes.ProductSubRoutes(app.Group("/product"))

	app.Logger.Fatal(app.Start(":1323"))
}
