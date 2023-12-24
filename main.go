package main

import (
	"github.com/labstack/echo/v4"
	"github.com/norbusonam/twitter-with-go-and-htmx/pkg/handlers"
)

func main() {
	e := echo.New()

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
	e.POST("/api/post/:id/reply", handlers.CreateReplyPost)
	e.DELETE("/api/post/:id", handlers.DeletePost)
	e.POST("/api/post/:id/like", handlers.LikePost)
	e.DELETE("/api/post/:id/like", handlers.UnlikePost)

	e.Logger.Fatal(e.Start(":8080"))
}
