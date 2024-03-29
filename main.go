package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/norbusonam/twitter-with-go-and-htmx/pkg/db"
	"github.com/norbusonam/twitter-with-go-and-htmx/pkg/handlers"
)

func main() {
	e := echo.New()

	// +--------------------+
	// | Logging Middleware |
	// +--------------------+
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod: true,
		LogStatus: true,
		LogURI:    true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			fmt.Printf("REQUEST: method: %v, uri: %v, status: %v\n", v.Method, v.URI, v.Status)
			return nil
		},
	}))

	// +--------------+
	// | Static Files |
	// +--------------+
	e.Static("/", "public")

	// +-------------+
	// | Page Routes |
	// +-------------+
	e.GET("/", handlers.Root)
	e.GET("/home", handlers.HomePage)
	e.GET("/u/:username", handlers.UserPage)
	e.GET("/u/:username/likes", handlers.UserLikesPage)
	e.GET("/u/:username/following", handlers.UserFollowingPage)
	e.GET("/u/:username/followers", handlers.UserFollowersPage)
	e.GET("/p/:id", handlers.PostPage)

	// +---------------+
	// | Action Routes |
	// +---------------+
	e.POST("/signin", handlers.SignIn)
	e.POST("/signup", handlers.SignUp)
	e.POST("/signout", handlers.SignOut)
	e.POST("/post", handlers.CreatePost)
	e.DELETE("/post/:id", handlers.DeletePost)
	e.POST("/post/:id/like", handlers.LikePost)
	e.DELETE("/post/:id/like", handlers.UnlikePost)

	// +---------------------+
	// | Initialize Database |
	// +---------------------+
	db.Init("postgres://postgres:postgres@localhost:5432/twitter?sslmode=disable")

	e.Logger.Fatal(e.Start(":8080"))
}
  