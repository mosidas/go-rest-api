package main

import (
	"net/http"

	"restapi/controller/authController"
	"restapi/controller/usersController"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", ping)
	e.GET("/index", index)
	// Users
	e.GET("/users/:id", usersController.GetById)
	e.GET("/users", usersController.GetByQuery)
	e.PUT("/users", usersController.Save)
	e.POST("/users", usersController.Save2)

	// Authentication
	e.GET("/auth", authController.SignIn)
	e.POST("auth", authController.SignUp)

	e.Logger.Fatal(e.Start("localhost:1323"))
}

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func index(c echo.Context) error {
	return c.File("./pages/index.html")
}
