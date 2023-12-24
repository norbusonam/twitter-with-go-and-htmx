package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SignIn(c echo.Context) error {
	return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented", c.Path()))
}

func SignUp(c echo.Context) error {
	return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented", c.Path()))
}

func SignOut(c echo.Context) error {
	return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented", c.Path()))
}
