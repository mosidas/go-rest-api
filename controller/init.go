package controller

import (
	"restapi/controller/auth"
	"restapi/controller/index"
	"restapi/controller/users"

	"github.com/labstack/echo/v4"
)

func AddController(e *echo.Echo) {
	index.Bind(e)
	users.Bind(e)
	auth.Bind(e)
}
