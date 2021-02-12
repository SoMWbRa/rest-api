package post

import (
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
	return c.String(http.StatusOK, "Return all posts")
}

func GetPost(c echo.Context) error {
	return c.String(http.StatusOK, "Return a single post")
}

func AddPost(c echo.Context) error {
	return c.String(http.StatusOK, "Adds  a new post")
}

func DeletePost(c echo.Context) error {
	return c.String(http.StatusOK, "Remove a post")
}
