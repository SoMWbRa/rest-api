package oauth

import (
	"fmt"
	"github.com/SoMWbRa/rest-api/database"
	"github.com/labstack/echo"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
)

var (
	googleOauthConfig *oauth2.Config
	oauthStateString  = "pseudo"
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

func HandleGoogleLogin(c echo.Context) error {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func HandleGoogleCallback(c echo.Context) error {
	content, err := getUserInfo(c.QueryParam("state"), c.QueryParam("code"))
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	return c.String(http.StatusOK, string(content))
}

func getUserInfo(state string, code string) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}
func InitConfig() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:3000/profile",
		ClientID:     "781063362244-qd1f7c491g8i9o46cudbad2seetfiomo.apps.googleusercontent.com",
		ClientSecret: "n0Ll7MbGYltnfJ7hgkKT_sy4",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}
