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
		fmt.Println("unable to find user")
		// TODO: handle error properly
		return err
	}
	if user.Password != password {
		fmt.Println("wrong password")
		// TODO: handle error properly
		return fmt.Errorf("wrong password")
	}
	c.SetCookie(&http.Cookie{
		Name:  "user_id",
		Value: strconv.Itoa(user.ID),
	})
	c.Response().Header().Set("HX-Location", "/home")
	return nil
}

func SignUp(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")
	user, err := db.CreateUser(username, email, password)
	if err != nil {
		fmt.Println("unable to create user")
		// TODO: handle error properly
		return err
	}
	c.SetCookie(&http.Cookie{
		Name:  "user_id",
		Value: strconv.Itoa(user.ID),
	})
	c.Response().Header().Set("HX-Location", "/home")
	return nil
}

func SignOut(c echo.Context) error {
	return c.String(http.StatusNotImplemented, fmt.Sprintf("%s not implemented", c.Path()))
}
