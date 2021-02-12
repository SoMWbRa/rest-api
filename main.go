package main

import (
	"fmt"
	"github.com/SoMWbRa/rest-api/database"
	"github.com/SoMWbRa/rest-api/post"
	"github.com/labstack/echo"
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
func main() {
	e := echo.New()

	initDatabase()

	e.GET("/api/v1/post", post.GetPosts)
	e.GET("/api/v1/post/:id", post.GetPost)
	e.POST("/api/v1/post", post.AddPost)
	e.DELETE("/api/v1/post/:id", post.DeletePost)

	e.Logger.Fatal(e.Start(":3000"))
}
