package review

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/foundation-13/gpr/pkg/types"
)

func Assemble(e *echo.Echo, m Manager) {
	h := &handler{
		manager: m,
	}

	g := e.Group("/reviews")

	g.POST("/", h.Create)
}

// impl

type handler struct {
	manager Manager
}

func (h *handler) Create(c echo.Context) error {
	ctx := c.Request().Context()

	var dto types.ReviewDTO
	err := c.Bind(&dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	id, err := h.manager.CreateReview(ctx, dto)
	if err != nil {
		return echo.NewHTTPError(500, err) // TODO: fix error handling
	}

	return c.JSON(http.StatusCreated, map[string]string{"id": id})
}
