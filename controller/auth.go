package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"restapi/model"
	"restapi/repository"
)

func SignIn(c echo.Context) error {
	u := new(model.Credential)
	if err := c.Bind(u); err != nil {
		return err
	}
	// check if user exists
	user, err := repository.GetById(u.Id)
	if err != nil {
		return err
	}
	// check if password is correct
	if user.Auth.Password != u.Password {
		return c.String(http.StatusUnauthorized, "Wrong password or username.")
	}
	return c.JSON(http.StatusOK, user)
}

func SignUp(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}
