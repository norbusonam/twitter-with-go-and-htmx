package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/norbusonam/twitter-with-go-and-htmx/pkg/templates"
)

func Root(c echo.Context) error {
	if isUserAuthenticated(c.Cookies()) {
		return c.Redirect(302, "/home")
	}
	return templates.Landing().Render(c.Request().Context(), c.Response().Writer)
}
