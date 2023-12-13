package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

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

	e.Logger.Fatal(e.Start(":8080"))
}
