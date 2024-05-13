package middleware

import (
	"github.com/gofiber/fiber/v2"
)

type mw struct {
}

func NewMW() MW {
	return &mw{}
}

type Cookie struct {
	UserId int64 `cookie:"user_id"`
}

// TODO: add jwt logic
func (m *mw) ValidateSession() fiber.Handler {
	return func(c *fiber.Ctx) error {
		cook := new(Cookie)
		err := c.CookieParser(cook)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).Render("sign_in", fiber.Map{})
		}

		// set in context
		c.Locals("user_id", cook.UserId)

		// TODO: inster session in context
		return c.Next()
	}
}
