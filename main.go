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
		// if user is logged in
		//  redirect to /home
		// else
		//  render login/home page
		return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented", c.Path()))
	})

	e.GET("/home", func(c echo.Context) error {
		// if user is logged in
		//  render home page
		// else
		//  redirect to /
		return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented", c.Path()))
	})

	e.GET("/explore", func(c echo.Context) error {
		// if user is logged in
		//  render explore page
		// else
		//  redirect to /
		return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented", c.Path()))
	})

	e.GET("/bookmarks", func(c echo.Context) error {
		// if user is logged in
		//  render bookmarks page
		// else
		//  redirect to /
		return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented", c.Path()))
	})

	e.GET("/u/:username", func(c echo.Context) error {
		// render user page
		username := c.Param("username")
		return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented (username=%s)", c.Path(), username))
	})

	e.GET("/p/:id", func(c echo.Context) error {
		// render post page
		postId := c.Param("id")
		return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented (postId=%s)", c.Path(), postId))
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

	e.POST("/api/post/:id/reply", func(c echo.Context) error {
		postId := c.Param("id")
		return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented (postId=%s)", c.Path(), postId))
	})

	e.DELETE("/api/post/:id", func(c echo.Context) error {
		postId := c.Param("id")
		return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented (postId=%s)", c.Path(), postId))
	})

	e.POST("/api/post/:id/like", func(c echo.Context) error {
		postId := c.Param("id")
		return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented (postId=%s)", c.Path(), postId))
	})

	e.DELETE("/api/post/:id/like", func(c echo.Context) error {
		postId := c.Param("id")
		return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented (postId=%s)", c.Path(), postId))
	})

	e.Logger.Fatal(e.Start(":8080"))
}
