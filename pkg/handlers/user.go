package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/norbusonam/twitter-with-go-and-htmx/pkg/db"
	"github.com/norbusonam/twitter-with-go-and-htmx/pkg/templates"
)

func UserPage(c echo.Context) error {
	// render user page
	username := c.Param("username")
	user, err := db.GetUserByUsername(username)
	if err != nil {
		return err
	}
	authenticatedUser, err := getAuthenticatedUser(c.Cookies())
	if err != nil {
		return err
	}
	posts, err := db.GetPostsByUser(username)
	if err != nil {
		return err
	}
	templates.User(user, posts, authenticatedUser).Render(c.Request().Context(), c.Response().Writer)
	return nil
}

func UserLikesPage(c echo.Context) error {
	// render user likes page
	username := c.Param("username")
	return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented (username=%s)", c.Path(), username))
}

func UserFollowingPage(c echo.Context) error {
	// render user following page
	username := c.Param("username")
	return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented (username=%s)", c.Path(), username))
}

func UserFollowersPage(c echo.Context) error {
	// render user followers page
	username := c.Param("username")
	return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented (username=%s)", c.Path(), username))
}
