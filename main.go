package main

import (
	"fmt"
	"github.com/SoMWbRa/rest-api/database"
	_ "github.com/SoMWbRa/rest-api/docs"
	"github.com/SoMWbRa/rest-api/post"
	"github.com/labstack/echo"
	"github.com/swaggo/echo-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Database connection successfully opened")

	err = db.AutoMigrate(&post.Post{})
	if err != nil {
		panic("failed to migrated")
	}
	fmt.Println("Database Migrated")
	database.DB = db
}

// @title REST API ECHO
func main() {
	e := echo.New()

	initDatabase()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := e.Group("/api/v1")
	{
		v1.GET("/post", post.GetPosts)
		v1.GET("/post/:id", post.GetPost)
		v1.POST("/post", post.AddPost)
		v1.PUT("/post/:id", post.PutPost)
		v1.DELETE("/post/:id", post.DeletePost)
	}
	e.Logger.Fatal(e.Start(":3000"))
}
