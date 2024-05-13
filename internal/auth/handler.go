package auth

import "github.com/gofiber/fiber/v2"

type Handler interface {
	IndexView(c *fiber.Ctx) error
	SignIn() fiber.Handler
}
