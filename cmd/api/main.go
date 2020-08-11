package main

import (
	"net/http"

	"github.com/labstack/echo"
	echomiddleware "github.com/labstack/echo/middleware"

	"github.com/foundation-13/gpr/pkg/api/middleware"
	"github.com/foundation-13/gpr/pkg/api/profile"
	"github.com/foundation-13/gpr/pkg/api/review"
	"github.com/foundation-13/gpr/pkg/log"
)

func main() {
	log.InitLog(true)

	e := echo.New()

	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())

	verifier := middleware.NewDummyTokenVerifier()
	authMdlwr := middleware.NewAuth(verifier)

	e.Use(authMdlwr.MiddlewareFunc)

	log.L.Info("api started")

	reviewManager := review.NewManager()
	review.Assemble(e, reviewManager)

	userManager := profile.NewManager()
	profile.Assemble(e, userManager)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "green"})
	})

	e.Logger.Fatal(e.Start(":8765"))
}
