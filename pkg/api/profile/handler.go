package profile

import (
	"github.com/foundation-13/gpr/pkg/utils"
	"github.com/labstack/echo"
	"net/http"
)

func Assemble(e *echo.Echo, m Manager, middleware ...echo.MiddlewareFunc) {
	h := &handler{
		manager: m,
	}

	g := e.Group("/profile")
	g.Use(middleware...)

	g.GET("/reviews", h.reviews)
}

type handler struct {
	manager Manager
}

func (h *handler) reviews(c echo.Context) error {
	ctx, userID := utils.FromEchoContext(c)

	reviews, err := h.manager.GetReviews(ctx, userID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, reviews)
}