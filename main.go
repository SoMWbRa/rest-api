package main

import (
	"./database"
	"./post"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
)

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "post.db")
	if err != nil {
		panic("Failed to connect databases")
	}
	fmt.Println("Database connection successfully opened")

	database.DBConn.AutoMigrate(&post.Post{})
	fmt.Println("Database Migrated")
}
func main() {
	e := echo.New()

	initDatabase()
	defer database.DBConn.Close()

	e.GET("/api/v1/post", post.GetPosts)
	e.GET("/api/v1/post/:id", post.GetPost)
	e.POST("/api/v1/post", post.AddPost)
	e.DELETE("/api/v1/post/:id", post.DeletePost)

	e.Logger.Fatal(e.Start(":3000"))
}
