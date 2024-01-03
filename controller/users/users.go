package users

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"restapi/model"
)

func Bind(e *echo.Echo) {
	e.GET("/users/:id", GetById)
	e.GET("/users", GetByQuery)
	e.PUT("/users", Save)
	e.POST("/users", Save2)
}

func GetById(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func GetByQuery(c echo.Context) error {
	q := c.QueryParam("q")
	return c.String(http.StatusOK, q)
}

func Save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusCreated, name+" "+email)
}

func Save2(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, user)
}
