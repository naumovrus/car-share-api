package httpServer

import (
	"car-api/config"
	"car-api/pkg/logger"

	"github.com/gofiber/fiber/v2"
	//"github.com/gofiber/fiber/v2/middleware/recover"
)

func (s *Server) MapHandlers(app *fiber.App, logger *logger.ApiLogger, cfg *config.Config) error {

	return nil
}
