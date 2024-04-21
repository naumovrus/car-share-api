package middleware

import "github.com/gofiber/fiber/v2"

type MW interface {
	ValidateJWT() fiber.Handler
	CheckRole() fiber.Handler
}
