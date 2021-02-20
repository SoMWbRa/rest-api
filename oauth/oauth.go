package oauth

import (
	"github.com/SoMWbRa/rest-api/database"
	"github.com/labstack/echo"
	"net/http"
)

type User struct {
	ID       int    `gorm:"autoIncrement" json:"id" `
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUser godoc
// @Summary create new user
// @Tags User
// @Router /user/register [post]
func CreateUser(c echo.Context) error {
	var user User
	err := c.Bind(&user)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid body")
	}
	db := database.DB
	db.Create(&user)
	return c.JSON(http.StatusCreated, &user)
}

// GetUser godoc
// @Summary get user info
// @Tags User
// @Router /user/login [post]
func GetUser(c echo.Context) error {
	var user User
	err := c.Bind(&user)
	if err != nil {
		return err
	}
	var userFind User
	db := database.DB
	db.First(&userFind, "email = ?", user.Email)
	if userFind.Password != user.Password || userFind.Email == "" {
		return c.String(http.StatusBadRequest, "invalid email or password")
	}
	return c.JSON(http.StatusOK, &userFind)
}
