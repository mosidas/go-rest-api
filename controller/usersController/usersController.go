package usersController

import (
	"net/http"
	"restapi/model"

	"github.com/labstack/echo/v4"
)

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
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}
