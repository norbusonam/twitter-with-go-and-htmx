package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/norbusonam/twitter-with-go-and-htmx/pkg/templates"
)

func HomePage(c echo.Context) error {
	// if user is logged in
	//  render home page
	// else
	//  redirect to /
	return templates.Home().Render(c.Request().Context(), c.Response().Writer)
}
