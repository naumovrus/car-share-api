package httpErrorHandler

import (
	"fmt"
	"log"

	"car-api/config"
	"car-api/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type HttpErrorHandler struct {
	showUnknownErrors bool
	logger            logger.Logger
}

func NewErrorHandler(c *config.Config, logger logger.Logger) *HttpErrorHandler {
	return &HttpErrorHandler{
		showUnknownErrors: c.Server.ShowUnknownErrorsInResponse,
		logger:            logger,
	}
}

type responseMsg struct {
	Message string `json:"message"`
}

func (handler *HttpErrorHandler) Handler(c *fiber.Ctx, err error) error {
	var response responseMsg
	var statusCode int

	if statusCode == 0 {
		statusCode = fiber.StatusInternalServerError
	}
	if response.Message == "" && handler.showUnknownErrors {
		response.Message = fmt.Sprintf("Error: \n\n %s", err.Error())
	} else if response.Message == "" {
		handler.logger.Error(err)
		response.Message = "unknown error"
	}

	log.Print(err)

	return c.Status(statusCode).JSON(response)
}

func (handler *HttpErrorHandler) StackTraceHandler(c *fiber.Ctx, e interface{}) {
	if e == nil {
		e = "contact support"
	}
	handler.logger.ErrorFull(fmt.Errorf("%s %s: %+v", c.Path(), c.Method(), e))
	c.Status(500).JSON(map[string]interface{}{
		"description": e,
		"status":      "500",
	})
}
