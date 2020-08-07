package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/foundation-13/gpr/pkg/log"
)

func main() {
	log.InitLog(true)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	log.L.Info("api started")

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "green"})
	})

	e.Logger.Fatal(e.Start(":8765"))
}
