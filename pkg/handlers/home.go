package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/norbusonam/twitter-with-go-and-htmx/pkg/db"
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
	posts, err := db.GetPosts()
	if err != nil {
		return err
	}
	return templates.Home(user, posts).Render(c.Request().Context(), c.Response().Writer)
}
