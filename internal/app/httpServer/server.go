package httpServer

import (
	"car-api/config"
	"car-api/pkg/httpErrorHandler"
	"car-api/pkg/logger"
	"fmt"

	"github.com/gofiber/fiber/v2"

	gojson "github.com/goccy/go-json"
)

type Server struct {
	fiber     *fiber.App
	cfg       *config.Config
	apiLogger *logger.ApiLogger
}

func NewServer(cfg *config.Config, apiLogger *logger.ApiLogger, handler *httpErrorHandler.HttpErrorHandler, shield *secure.Shield) *Server {
	return &Server{
		fiber: fiber.New(fiber.Config{
			ErrorHandler:          handler.Handler,
			DisableStartupMessage: false,
			JSONEncoder:           gojson.Marshal,
			JSONDecoder:           gojson.Unmarshal,
		}),
		cfg:       cfg,
		apiLogger: apiLogger,
	}
}

func (s *Server) Run() error {
	if err := s.MapHandlers(s.fiber, s.apiLogger, s.cfg); err != nil {
		s.apiLogger.Fatalf("Cannot map handlers: %+v", err)
	}
	s.apiLogger.Infof("Start server on port: %s:%s", s.cfg.Server.Host, s.cfg.Server.Port)
	if err := s.fiber.Listen(fmt.Sprintf("%s:%s", s.cfg.Server.Host, s.cfg.Server.Port)); err != nil {
		s.apiLogger.Fatalf("Error starting Server: %+v", err)
	}

	return nil
}
