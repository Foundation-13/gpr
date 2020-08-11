package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"

	"github.com/foundation-13/gpr/pkg/utils"
)

const (
	authToken     = "Authorization"
	authTokenHead = "Bearer"
)

//go:generate mockery -name TokenVerifier -outpkg middlewaremocks -output ./middlewaremocks -dir .
type TokenVerifier interface {
	VerifyToken(token string) (string, error)
}

// Authentication middleware
// Extract bearer token from the headers and check it

type Auth struct {
	verifier TokenVerifier
}

func NewAuth(verifier TokenVerifier) *Auth {
	return &Auth{verifier: verifier}
}

func (m *Auth) MiddlewareFunc(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get(authToken)
		if len(authHeader) == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "Bearer token not found")
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 && parts[0] != authTokenHead {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid Bearer token")
		}

		token := parts[1]

		userID, err := m.verifier.VerifyToken(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token validating error")
		}

		utils.UpdateEchoContextWithUserID(c, userID)

		return next(c)
	}
}
