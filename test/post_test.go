package test

import (
	"github.com/SoMWbRa/rest-api/post"
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
	postJSON = `{"user_id":1,"id":1,"title":"t1","body":"b1"} `
	putJSON  = `{"user_id":1,"id":8,"title":"t8","body":"b8"} `
)

// TestAddPost doc
// Структура тестов, нужно инициализировать новую БД
func TestAddPost(t *testing.T) {
	e := server.InitiateServer()
	err := server.InitDatabase("test.db")
	if err != nil {
		log.Fatalln(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/api/v1/post", strings.NewReader(postJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, post.AddPost(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "{\"user_id\":1,\"id\":1,\"title\":\"t1\",\"body\":\"b1\"}\n", rec.Body.String())
	}
}

func TestGetPost(t *testing.T) {
	e := server.InitiateServer()
	err := server.InitDatabase("test.db")
	if err != nil {
		log.Fatalln(err)
	}
	req := httptest.NewRequest(http.MethodGet, "/api/v1/post?data=json", strings.NewReader(postJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, post.GetPost(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"user_id\":1,\"id\":1,\"title\":\"t1\",\"body\":\"b1\"}\n", rec.Body.String())
	}
}
func TestPutPost(t *testing.T) {
	e := server.InitiateServer()
	err := server.InitDatabase("test.db")
	if err != nil {
		log.Fatalln(err)
	}
	req := httptest.NewRequest(http.MethodPut, "/api/v1/post/1", strings.NewReader(putJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, post.PutPost(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"user_id\":1,\"id\":8,\"title\":\"t8\",\"body\":\"b8\"}\n", rec.Body.String())
	}
}

func TestDeletePost(t *testing.T) {
	e := server.InitiateServer()
	err := server.InitDatabase("test.db")
	if err != nil {
		log.Fatalln(err)
	}
	req := httptest.NewRequest(http.MethodGet, "/api/v1/post/1", strings.NewReader(postJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, post.DeletePost(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
