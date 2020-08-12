package review

import (
	"github.com/foundation-13/gpr/pkg/log"
	"net/http"

	"github.com/gavrilaf/errors"
	"github.com/labstack/echo"

	"github.com/foundation-13/gpr/pkg/types"
	"github.com/foundation-13/gpr/pkg/utils"
)

func Assemble(e *echo.Echo, m Manager, middleware ...echo.MiddlewareFunc) {
	h := &handler{
		manager: m,
	}

	g := e.Group("/reviews")
	g.Use(middleware...)

	g.POST("", h.Create)
	g.PUT("/upload", h.AddImage)
}

// impl

type handler struct {
	manager Manager
}

func (h *handler) Create(c echo.Context) error {
	var dto types.ReviewDTO
	err := c.Bind(&dto)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	ctx, userID := utils.FromEchoContext(c)

	id, err := h.manager.CreateReview(ctx, userID, dto)
	if err != nil {
		return echo.NewHTTPError(500, err) // TODO: fix error handling
	}

	return c.JSON(http.StatusCreated, map[string]string{"id": id})
}

func (h *handler) AddImage(c echo.Context) error {
	fileHeader, err := c.FormFile("file")
	ctx := c.Request().Context()
	if err != nil {
		err = errors.NewBadRequest(err, "file not found")
		log.L.WithError(err).Error("failed to add image into gallery")
		return err
	}

	file, err := fileHeader.Open()
	if err != nil {
		err = errors.NewBadRequest(err, "couldn't open file")
		log.L.WithError(err).Error("failed to add image into gallery")
		return err
	}
	defer file.Close()

	err = h.manager.AddImage(ctx,file, fileHeader.Filename, fileHeader.Header.Get("Content-Type"))
	if err != nil {
		log.L.WithError(err).Error("failed to add image into gallery")
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"success": "true"})
}