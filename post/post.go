package post

import (
	"github.com/SoMWbRa/rest-api/database"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"net/http"
)

type Post struct {
	UserId int    `json:"user_id" xml:"user_id"`
	ID     int    `json:"id" xml:"id"`
	Title  string `json:"title" xml:"title"`
	Body   string `json:"body" xml:"body"`
}

func GetPosts(c echo.Context) error {
	db := database.DB
	var posts []Post
	db.Find(&posts)
	data := c.QueryParam("data")
	if data == "json" {
		return c.JSONPretty(http.StatusOK, &posts, " ")
	} else if data == "xml" {
		return c.XMLPretty(http.StatusOK, &posts, " ")
	} else {
		return c.String(http.StatusBadRequest, "Select output format")
	}
}

func GetPost(c echo.Context) error {
	id := c.Param("id")
	db := database.DB
	var post Post
	db.Find(&post, id)
	data := c.QueryParam("data")
	if data == "json" {
		return c.JSONPretty(http.StatusOK, &post, " ")
	} else if data == "xml" {
		return c.XMLPretty(http.StatusOK, &post, " ")
	} else {
		return c.String(http.StatusBadRequest, "Select output format")
	}
}

func AddPost(c echo.Context) error {
	db := database.DB
	var post Post
	if err := c.Bind(&post); err != nil {
		return err
	}
	db.Create(&post)
	return c.JSONPretty(http.StatusCreated, &post, " ")
}

func DeletePost(c echo.Context) error {
	id := c.Param("id")
	db := database.DB
	var post Post
	db.First(&post, id)
	if post.ID == 0 {
		return c.String(http.StatusInternalServerError, "post doesn't exist")
	}
	db.Delete(&post, id)
	return c.String(http.StatusOK, "Post successfully deleted")
}
