package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/norbusonam/twitter-with-go-and-htmx/pkg/templates"
)

func HomePage(c echo.Context) error {
	user, err := getAuthenticatedUser(c.Cookies())
	if err != nil {
		return err
	}
	if user == nil {
		return c.Redirect(302, "/")
	}

	return templates.Home(user).Render(c.Request().Context(), c.Response().Writer)
}
