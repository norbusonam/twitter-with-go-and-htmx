package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/norbusonam/twitter-with-go-and-htmx/pkg/templates"
)

func Root(c echo.Context) error {
	// if user is logged in
	//  redirect to /home
	// else
	//  render login/home page
	return templates.Landing().Render(c.Request().Context(), c.Response().Writer)
}
