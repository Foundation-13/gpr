package main

import (
	"encoding/json"
	"fmt"
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

	log.L.Info("api started")

	reviewManager := review.NewManager()
	review.Assemble(e, reviewManager, authMdlwr.MiddlewareFunc)

	userManager := profile.NewManager()
	profile.Assemble(e, userManager, authMdlwr.MiddlewareFunc)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "green"})
	})


	data, _ := json.MarshalIndent(e.Routes(), "", "  ")

	fmt.Println(string(data))

	e.Logger.Fatal(e.Start(":8765"))
}
