package main

import (
	"car-api/config"
	"car-api/internal/app/httpServer"
	"car-api/pkg/httpErrorHandler"
	"car-api/pkg/logger"
	"log"

	"github.com/gofiber/template/html/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	log.Println("Starting server")
	v, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Cannot cload config: ", err.Error())
	}
	cfg, err := config.ParseConfig(v)
	if err != nil {
		log.Fatalf("Config parse error", err.Error())
	}
	log.Println("Config loaded")

	appLogger := logger.NewApiLogger(cfg)
	err = appLogger.InitLogger()
	if err != nil {
		log.Fatalf("Cannot init logger: %v", err.Error())
	}
	engine := html.New("./internal/templates", ".tmpl")
	appLogger.Infof("Logger successfully started with - Level: %s",
		cfg.Logger.Level)

	httpErrorHandler := httpErrorHandler.NewErrorHandler(cfg, appLogger)
	s := httpServer.NewServer(cfg, appLogger, httpErrorHandler, engine)
	if err = s.Run(); err != nil {
		appLogger.ErrorFull(err)
	}

}
