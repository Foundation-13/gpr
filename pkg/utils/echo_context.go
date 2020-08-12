package utils

import (
	"context"
	"github.com/labstack/echo"
)

const userIDKey = "user_id"

func UpdateEchoContextWithUserID(c echo.Context, userID string) {
	c.Set(userIDKey, userID)
}

func FromEchoContext(c echo.Context) (context.Context, string) {
	userID := c.Get(userIDKey)
	return c.Request().Context(), userID.(string)
}
