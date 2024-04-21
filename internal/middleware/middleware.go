package middleware

import "github.com/gofiber/fiber/v2"

type mw struct {
}

func NewMW() MW {
	return &mw{}
}

// TODO: add jwt logic
func (m *mw) ValidateJWT() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}

func (m *mw) CheckRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}
