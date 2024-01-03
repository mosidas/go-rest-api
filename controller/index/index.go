package index

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Bind(e *echo.Echo) {
	e.GET("/ping", ping)
	e.GET("/", index)
}

func ping(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func index(c echo.Context) error {
	return c.File("./pages/index.html")
}
