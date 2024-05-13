package middleware

import "github.com/gofiber/fiber/v2"

type MW interface {
	ValidateSession() fiber.Handler
}
