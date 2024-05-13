package httpServer

import (
	"car-api/config"
	"car-api/pkg/logger"
	storage_postgres "car-api/pkg/storage"

	authHttp "car-api/internal/auth/delivery/http"
	authRepo "car-api/internal/auth/repository"
	authUC "car-api/internal/auth/usecase"
	"car-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
	//"github.com/gofiber/fiber/v2/middleware/recover"
)

func (s *Server) MapHandlers(app *fiber.App, logger *logger.ApiLogger, cfg *config.Config) error {
	// Init services
	dbLocal, err := storage_postgres.InitPsqlDB(&s.cfg.DbLocal)
	if err != nil {
		return err
	}

	//init repos
	authRepo := authRepo.New(dbLocal)
	// orderRepo := orderRepo.New(dbLocal)
	mw := middleware.NewMW()
	// fhttpClient := fhttp.NewClient()

	//init usecases
	authUC := authUC.New(authRepo)

	//init handlers
	authHandlers := authHttp.New(authUC, *s.apiLogger)

	//init groups

	authGroup := app.Group("auth")

	app.Get("/", mw.ValidateSession(), authHandlers.IndexView)

	authHttp.MapAuthRoutes(authGroup, mw, authHandlers)

	return nil
}
