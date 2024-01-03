package auth

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"restapi/model"
	"restapi/repository"
)

func Bind(e *echo.Echo) {
	e.GET("/auth", SignIn)
	e.POST("/auth", SignUp)
}

func SignIn(c echo.Context) error {
	u := new(model.Credential)
	if err := c.Bind(u); err != nil {
		return err
	}
	// check if user exists
	var userRepo repository.IUserRepository = repository.NewUserRepository()
	println(userRepo)
	user, err := userRepo.GetById(u.Id)
	if err != nil {
		return err
	}
	// check if password is correct
	if user.Auth.Password != u.Password {
		return c.String(http.StatusUnauthorized, "Wrong password or username.")
	}

	// セッションCookieの作成
	cookie := new(http.Cookie)
	cookie.Name = "session_id"
	cookie.Value = "some_session_id"
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteLaxMode
	cookie.Secure = true

	// Cookieをセット
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, user)
}

func SignUp(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}
