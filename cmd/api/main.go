package main

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
	echomiddleware "github.com/labstack/echo/middleware"

	"github.com/foundation-13/gpr/pkg/api/middleware"
	"github.com/foundation-13/gpr/pkg/api/profile"
	"github.com/foundation-13/gpr/pkg/api/review"
	"github.com/foundation-13/gpr/pkg/log"
	"github.com/foundation-13/gpr/pkg/storage"
	"github.com/foundation-13/gpr/pkg/utils"
)

func main() {
	log.InitLog(true)

	e := echo.New()

	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())

	verifier := middleware.NewDummyTokenVerifier()
	authMdlwr := middleware.NewAuth(verifier)
	ctx := context.Background()
	bucket, err := storage.NewGCPBucket(ctx, "image-test-bucket-123")
	if err != nil {
		log.L.WithError(err).Panic("failed to connect bucket")
	}

	reviewConfig := review.Config{
		Storage: bucket,
		IDGen: utils.UUIDIDGen{},
	}

	log.L.Info("api started")

	reviewManager := review.NewManager(reviewConfig)
	review.Assemble(e, reviewManager, authMdlwr.MiddlewareFunc)


	userManager := profile.NewManager()
	profile.Assemble(e, userManager, authMdlwr.MiddlewareFunc)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "green"})
	})

	e.Logger.Fatal(e.Start(":8765"))
}
