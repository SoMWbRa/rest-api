package test

import (
	"github.com/SoMWbRa/rest-api/oauth"
	"github.com/SoMWbRa/rest-api/server"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	createJSON = `{"id":1,"email":"test@gmail.com","password":"password"}`
)

func TestCreateUser(t *testing.T) {
	e := server.InitiateServer()
	err := server.InitDatabase("test.db")
	if err != nil {
		log.Fatalln(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/user/register", strings.NewReader(createJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, oauth.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "{\"id\":1,\"email\":\"test@gmail.com\",\"password\":\"password\"}\n", rec.Body.String())
	}
}
func GetUser(t *testing.T) {
	e := server.InitiateServer()
	err := server.InitDatabase("test.db")
	if err != nil {
		log.Fatalln(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/user/login", strings.NewReader(createJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, oauth.GetUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"id\":1,\"email\":\"test@gmail.com\",\"password\":\"password\"}\n", rec.Body.String())
	}
}
