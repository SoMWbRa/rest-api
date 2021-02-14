package test

import (
	"github.com/SoMWbRa/rest-api/post"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	postJSON = `{"user_id":1,"id":1,"title":"t1","body":"b1"}`
)

// TestAddPost doc
// Структура тестов, нужно инициализировать новую БД
func TestAddPost(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(postJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, post.AddPost(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, postJSON, rec.Body.String())
	}

}
