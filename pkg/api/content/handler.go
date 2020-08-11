package content

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/foundation-13/gpr/pkg/types"
)

func Assemble(e *echo.Echo, m Manager) {
	h := &handler{
		manager: m,
	}

	g := e.Group("/content")

	g.POST("/create", h.create)
}

type handler struct {
	manager Manager
}

func(h *handler) create(c echo.Context) (err error){
	dto :=new(types.ReviewDTO)

	if err = c.Bind(dto); err != nil{
		return
	}
	return c.JSON(http.StatusOK, map[string]string{"info":dto.Info, "stars":dto.Stars})
}
