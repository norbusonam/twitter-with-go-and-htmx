package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// +--------------------+
	// | User Facing Routes |
	// +--------------------+
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented", c.Path()))
	})

	e.GET("/home", func(c echo.Context) error {
		return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented", c.Path()))
	})

	e.GET("/explore", func(c echo.Context) error {
		return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented", c.Path()))
	})

	e.GET("/bookmarks", func(c echo.Context) error {
		return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented", c.Path()))
	})

	e.GET("/:username", func(c echo.Context) error {
		username := c.Param("username")
		return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented (username=%s)", c.Path(), username))
	})

	e.GET("/:username/:postId", func(c echo.Context) error {
		username := c.Param("username")
		postId := c.Param("postId")
		return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented (username=%s, postId=%s)", c.Path(), username, postId))
	})

	// +-----------------------+
	// | Service Facing Routes |
	// +-----------------------+

	e.POST("/api/signIn", func(c echo.Context) error {
		return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented", c.Path()))
	})

	e.POST("/api/signUp", func(c echo.Context) error {
		return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented", c.Path()))
	})

	e.POST("/api/signOut", func(c echo.Context) error {
		return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented", c.Path()))
	})

	e.POST("/api/post", func(c echo.Context) error {
		return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented", c.Path()))
	})

	e.POST("/api/post/:postId/reply", func(c echo.Context) error {
		postId := c.Param("postId")
		return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented (postId=%s)", c.Path(), postId))
	})

	e.DELETE("/api/post/:postId", func(c echo.Context) error {
		postId := c.Param("postId")
		return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented (postId=%s)", c.Path(), postId))
	})

	e.POST("/api/post/:postId/like", func(c echo.Context) error {
		postId := c.Param("postId")
		return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented (postId=%s)", c.Path(), postId))
	})

	e.DELETE("/api/post/:postId/like", func(c echo.Context) error {
		postId := c.Param("postId")
		return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented (postId=%s)", c.Path(), postId))
	})

	e.Logger.Fatal(e.Start(":8080"))
}
