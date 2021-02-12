package post

import (
	"github.com/SoMWbRa/rest-api/database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"net/http"
)

type Post struct {
	gorm.Model
	UserId int    `json:"user_id"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func GetPosts(c echo.Context) error {
	db := database.DBConn
	var posts []Post
	db.Find(&posts)
	return c.JSON(http.StatusOK, &posts)
}

func GetPost(c echo.Context) error {
	id := c.Param("id")
	db := database.DBConn
	var post Post
	db.Find(&post, id)
	return c.JSON(http.StatusOK, &post)
}

func AddPost(c echo.Context) error {
	db := database.DBConn
	var post Post
	post.UserId = 1
	post.Id = 1
	post.Title = "t1"
	post.Body = "b1"

	db.Create(&post)
	return c.JSON(http.StatusOK, &post)
}

func DeletePost(c echo.Context) error {
	return c.String(http.StatusOK, "Remove a post")
}
