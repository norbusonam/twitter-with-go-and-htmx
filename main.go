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
		LogStatus: true,
		LogURI:    true,
		BeforeNextFunc: func(c echo.Context) {
			c.Set("customValueFromContext", 42)
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			value, _ := c.Get("customValueFromContext").(int)
			fmt.Printf("REQUEST: uri: %v, status: %v, custom-value: %v\n", v.URI, v.Status, value)
			return nil
		},
	}))

	// +--------------+
	// | Static Files |
	// +--------------+
	e.Static("/", "public")

	// +--------------------+
	// | User Facing Routes |
	// +--------------------+
	e.GET("/", handlers.Root)
	e.GET("/home", handlers.HomePage)
	e.GET("/u/:username", handlers.UserPage)
	e.GET("/u/:username/likes", handlers.UserLikesPage)
	e.GET("/u/:username/following", handlers.UserFollowingPage)
	e.GET("/u/:username/followers", handlers.UserFollowersPage)
	e.GET("/p/:id", handlers.PostPage)

	// +-----------------------+
	// | Service Facing Routes |
	// +-----------------------+
	e.POST("/api/signin", handlers.SignIn)
	e.POST("/api/signup", handlers.SignUp)
	e.POST("/api/signout", handlers.SignOut)
	e.POST("/api/post", handlers.CreatePost)
	e.DELETE("/api/post/:id", handlers.DeletePost)
	e.POST("/api/post/:id/like", handlers.LikePost)
	e.DELETE("/api/post/:id/like", handlers.UnlikePost)

	// +---------------------+
	// | Initialize Database |
	// +---------------------+
	db.Init("postgres://postgres:postgres@localhost:5432/twitter?sslmode=disable")

	e.Logger.Fatal(e.Start(":8080"))
}
