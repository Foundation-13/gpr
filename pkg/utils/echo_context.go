package utils

import (
	"github.com/labstack/echo"
)

const userIDKey = "user_id"

func UpdateEchoContextWithUserID(c echo.Context, userID string) {
	c.Set(userIDKey, userID)
}

func UserIDFromEchoContext(c echo.Context) string {
	userID := c.Get(userIDKey)
	return userID.(string)
}
