package mdlwr_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	. "github.com/foundation-13/gpr/pkg/api/mdlwr"
)

func TestMdlwrAuth(t *testing.T) {
	tests := []struct {
		name   string
		header string
		code   int
	}{
		{"no token", "", http.StatusBadRequest},
		{"invalid token", "Bearrrrer", http.StatusBadRequest},
		{"expired token", "Bearer invalid-token", http.StatusUnauthorized},
		{"valid token", FakeValidBearerAuthToken, http.StatusOK},
	}

	mdlwr := NewAuthMdlwr(NewFakeAuthTokenVerifier())

	e := echo.New()
	e.GET("/token", mdlwr.MiddlewareFunc(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	}))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			req, _ := http.NewRequest("GET", "/token", nil)
			if len(tt.header) > 0 {
				req.Header.Set("Authorization", tt.header)
			}
			e.ServeHTTP(w, req)

			assert.Equal(t, tt.code, w.Code)
		})
	}
}