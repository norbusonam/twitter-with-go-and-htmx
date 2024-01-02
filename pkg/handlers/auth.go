package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/norbusonam/twitter-with-go-and-htmx/pkg/db"
)

func SignIn(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	user, err := db.GetUserByUsername(username)
	if err != nil {
		return err
	}
	if user.Password != password {
		return fmt.Errorf("invalid password")
	}
	c.SetCookie(&http.Cookie{
		Name:  "user_id",
		Value: strconv.Itoa(user.ID),
	})
	return c.Redirect(http.StatusFound, "/home")
}

func SignUp(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")
	fmt.Println(username, email, password)
	user, err := db.CreateUser(username, email, password)
	if err != nil {
		return err
	}
	c.SetCookie(&http.Cookie{
		Name:  "user_id",
		Value: strconv.Itoa(user.ID),
	})
	return c.Redirect(http.StatusFound, "/home")
}

func SignOut(c echo.Context) error {
	return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented", c.Path()))
}
