package middleware_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/foundation-13/gpr/pkg/api/middleware"
	"github.com/foundation-13/gpr/pkg/api/middleware/middlewaremocks"
)

func TestAuth(t *testing.T) {
	tests := []struct {
		name         string
		header       string
		userID       string
		err          error
		expectedCode int
	}{
		{
			name:         "no token",
			header:       "",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "invalid token",
			header:       "Bearrrrer",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "validation error",
			header:       "Bearer 456",
			err:          fmt.Errorf(""),
			expectedCode: http.StatusUnauthorized,
		},
		{
			name:         "valid token",
			header:       "Bearer 123",
			userID:       "1",
			expectedCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			verifier := &middlewaremocks.TokenVerifier{}
			m := middleware.NewAuth(verifier)

			e := echo.New()
			e.GET("/token", m.MiddlewareFunc(func(c echo.Context) error {
				return c.String(http.StatusOK, "test")
			}))

			w := httptest.NewRecorder()

			verifier.On("VerifyToken", mock.Anything).Return(tt.userID, tt.err)

			req, _ := http.NewRequest("GET", "/token", nil)
			if len(tt.header) > 0 {
				req.Header.Set("Authorization", tt.header)
			}
			e.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
		})
	}
}
