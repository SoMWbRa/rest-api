package server

import (
	"github.com/SoMWbRa/rest-api/database"
	"github.com/SoMWbRa/rest-api/oauth"
	"github.com/SoMWbRa/rest-api/post"
	"github.com/labstack/echo"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitiateServer() *echo.Echo {
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := e.Group("/api/v1")
	{
		v1.GET("/post", post.GetPosts)
		v1.GET("/post/:id", post.GetPost)
		v1.POST("/post", post.AddPost)
		v1.PUT("/post/:id", post.PutPost)
		v1.DELETE("/post/:id", post.DeletePost)
	}

	e.POST("/user/register", oauth.CreateUser)
	e.POST("/user/login", oauth.GetUser)

	e.File("/", "login.html")
	e.GET("/login", oauth.HandleGoogleLogin)
	e.GET("/profile", oauth.HandleGoogleCallback)
	return e
}
func InitDatabase(base string) error {
	var err error
	db, err := gorm.Open(sqlite.Open(base), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&post.Post{}, &oauth.User{})
	if err != nil {
		return err
	}
	database.DB = db
	return nil
}
