package profile

import (
	"github.com/labstack/echo"
	"net/http"
)

func Assemble(e *echo.Echo, man Manager) {
	h := &handler{
		man: man,
	}

	g := e.Group("/profile")

	g.GET("/reviews", h.reviews)
}

type handler struct {
	man Manager
}

func (h *handler) reviews(c echo.Context) error {
	ctx := c.Request().Context()
	userID := "from ctx"

	reviews, err := h.man.GetReviews(ctx, userID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, reviews)
}