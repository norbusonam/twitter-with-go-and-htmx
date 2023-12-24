package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomePage(c echo.Context) error {
	// if user is logged in
	//  render home page
	// else
	//  redirect to /
	return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented", c.Path()))
}
