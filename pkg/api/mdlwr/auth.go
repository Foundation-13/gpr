package mdlwr

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"

	"github.com/foundation-13/gpr/pkg/api/utils"
)

const (
	authToken     = "Authorization"
	authTokenHead = "Bearer"
)

type AuthMdlwrTokenVerifier interface {
	VerifyToken(idToken string) (string, error)
}

type AuthMdlwr struct {
	verifier AuthMdlwrTokenVerifier
}

func NewAuthMdlwr(verifier AuthMdlwrTokenVerifier) *AuthMdlwr {
	return &AuthMdlwr{verifier: verifier}
}

// Authentication middleware
// Extract bearer token from the headers and check it

func (mdlwr *AuthMdlwr) MiddlewareFunc(next echo.HandlerFunc) echo.HandlerFunc {
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

		userID, err := mdlwr.verifier.VerifyToken(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token validating error")
		}

		utils.SetUserIdInContext(c, userID)
		return next(c)
	}
}