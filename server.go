package main

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"restapi/controller"
	"restapi/middleware/apikey"
)

func main() {
	e := echo.New()
	e.Debug = true
	e.Logger.SetLevel(1)

	controller.AddController(e)

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
