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

// GetPosts godoc
// @Summary return list of posts
// @Tags Posts
// @Router /api/v1/post [get]
func GetPosts(c echo.Context) error {
	db := database.DB
	var posts []Post
	db.Find(&posts)
	data := c.QueryParam("data")
	if data == "json" {
		return c.JSON(http.StatusOK, &posts)
	} else if data == "xml" {
		return c.XML(http.StatusOK, &posts)
	} else {
		return c.String(http.StatusBadRequest, "Select output format")
	}
}

// GetPost godoc
// @Summary  return post
// @Tags Posts
// @Router /api/v1/post/{id} [get]
func GetPost(c echo.Context) error {
	id := c.Param("id")
	db := database.DB
	var post Post
	db.Find(&post, id)
	data := c.QueryParam("data")
	if data == "json" {
		return c.JSON(http.StatusOK, &post)
	} else if data == "xml" {
		return c.XML(http.StatusOK, &post)
	} else {
		return c.String(http.StatusBadRequest, "Select output format")
	}
}

// AddPost godoc
// @Summary add post
// @Tags Posts
// @Router /api/v1/post/{id} [post]
func AddPost(c echo.Context) error {
	db := database.DB
	var post Post
	if err := c.Bind(&post); err != nil {
		return err
	}
	db.Create(&post)
	return c.JSON(http.StatusCreated, &post)
}

// PutPost godoc
// @Summary update user_id , title , body
// @Tags Posts
// @Router /api/v1/post/{id} [put]
func PutPost(c echo.Context) error {
	db := database.DB
	id := c.Param("id")
	var put Post
	err := c.Bind(&put)
	if err != nil {
		return err
	}
	var old Post
	db.First(&old, id)

	db.Model(&old).Update("user_id", put.UserId)
	db.Model(&old).Update("title", put.Title)
	db.Model(&old).Update("body", put.Body)

	db.First(&put, id)
	return c.JSON(http.StatusOK, &put)
}

// DeletePost godoc
// @Summary delete post
// @Tags Posts
// @Produce  json
// @Router /api/v1/post/{id} [delete]
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
