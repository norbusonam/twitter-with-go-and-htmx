package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func PostPage(c echo.Context) error {
	// render post page
	postId := c.Param("id")
	return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented (postId=%s)", c.Path(), postId))
}

func CreatePost(c echo.Context) error {
	return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented", c.Path()))
}

func CreateReplyPost(c echo.Context) error {
	postId := c.Param("id")
	return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented (postId=%s)", c.Path(), postId))
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
