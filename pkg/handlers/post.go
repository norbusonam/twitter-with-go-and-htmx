package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/norbusonam/twitter-with-go-and-htmx/pkg/db"
	"github.com/norbusonam/twitter-with-go-and-htmx/pkg/templates"
)

func PostPage(c echo.Context) error {
	// render post page
	postId := c.Param("id")
	return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented (postId=%s)", c.Path(), postId))
}

func CreatePost(c echo.Context) error {
	content := c.FormValue("content")
	user, err := getAuthenticatedUser(c.Cookies())
	if err != nil {
		return err
	}
	if user == nil {
		return c.Redirect(302, "/")
	}
	post, err := db.CreatePost(content, user.ID)
	if err != nil {
		return err
	}
	return templates.Post(post).Render(c.Request().Context(), c.Response().Writer)
}

func DeletePost(c echo.Context) error {
	postId := c.Param("id")
	return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented (postId=%s)", c.Path(), postId))
}

func LikePost(c echo.Context) error {
	postId := c.Param("id")
	return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented (postId=%s)", c.Path(), postId))
}

func UnlikePost(c echo.Context) error {
	postId := c.Param("id")
	return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented (postId=%s)", c.Path(), postId))
}
