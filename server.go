package main

import (
	"restapi/controller"
	"restapi/middleware/apikey"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Debug = true
	e.Logger.SetLevel(1)

	controller.AddController(e)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.PATCH},
	}))
	e.Use(apikey.AuthApiKey)
	e.Use(session.Middleware(sessions.NewFilesystemStore("", []byte("secret"))))

	e.Logger.Fatal(e.Start("localhost:1323"))
}
