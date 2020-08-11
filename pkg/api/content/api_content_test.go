package content

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var newReviewJson = `{"info":"info","stars":"5"}
`
var newInvalidJson = `{invalid}`

var errStr = "code=400, message=Syntax error: offset=2, error=invalid character 'i' looking for beginning of object key string"

func TestCreateReview(t *testing.T) {
	t.Run("succeeded", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/content/create", strings.NewReader(newReviewJson))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		m := NewManager()
		h := &handler{m}

		if assert.NoError(t, h.create(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, newReviewJson, rec.Body.String())
		}
	})
	t.Run("Invalid JSON", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/content/create", strings.NewReader(newInvalidJson))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		m := NewManager()
		h := &handler{m}

		assert.EqualError(t, h.create(c), errStr)

	})
}
