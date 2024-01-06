package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/norbusonam/twitter-with-go-and-htmx/pkg/db"
	"github.com/norbusonam/twitter-with-go-and-htmx/pkg/templates"
)

func SignIn(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	user, err := db.GetUserByUsername(username)
	if err != nil {
		// TODO: figure out how to set status code without superfluous call warning
		templates.AuthError("user not found").Render(c.Request().Context(), c.Response().Writer)
		return nil
	}
	if user.Password != password {
		// TODO: figure out how to set status code without superfluous call warning
		templates.AuthError("wrong password").Render(c.Request().Context(), c.Response().Writer)
		return nil
	}
	c.SetCookie(&http.Cookie{
		Name:  "user_id",
		Value: strconv.Itoa(user.ID),
	})
	c.Response().Header().Set("HX-Location", "/home")
	return nil
}

func SignUp(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")
	user, err := db.CreateUser(username, email, password)
	if err != nil {
		// TODO: figure out how to set status code without superfluous call warning
		templates.AuthError("unable to create user").Render(c.Request().Context(), c.Response().Writer)
		return nil
	}
	c.SetCookie(&http.Cookie{
		Name:  "user_id",
		Value: strconv.Itoa(user.ID),
	})
	c.Response().Header().Set("HX-Location", "/home")
	return nil
}

func SignOut(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:   "user_id",
		MaxAge: -1,
	})
	c.Response().Header().Set("HX-Location", "/")
	return nil
}
