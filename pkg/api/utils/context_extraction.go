package utils

import (
	"github.com/labstack/echo"
)

const userIDKey = "user_id"

func SetUserIdInContext(c echo.Context, userID string) {
	c.Set(userIDKey, userID)
}

func GetUserIdFromContext(c echo.Context) string {
	userID := c.Get(userIDKey)
	return userID.(string)
}
