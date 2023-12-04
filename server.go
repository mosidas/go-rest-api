package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"restapi/controller"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", ping)
	e.GET("/index", index)
	// Users
	e.GET("/users/:id", controller.GetById)
	e.GET("/users", controller.GetByQuery)
	e.PUT("/users", controller.Save)
	e.POST("/users", controller.Save2)

	// Authentication
	e.GET("/auth", controller.SignIn)
	e.POST("auth", controller.SignUp)

	e.Logger.Fatal(e.Start("localhost:1323"))
}

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func index(c echo.Context) error {
	return c.File("./pages/index.html")
}
