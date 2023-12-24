package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserPage(c echo.Context) error {
	// render user page
	username := c.Param("username")
	return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented (username=%s)", c.Path(), username))
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
