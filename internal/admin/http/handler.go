package http

import (
	"car-api/internal/admin"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
}

func New() admin.Handler {
	return &handler{}
}

func (r *handler) AdminView() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Render("index", fiber.Map{
			"IsAdmin":         isAdmin,
			"ClientId":        clientId,
			"IsAuthenticated": isAuthenticated,
		})
	}
}
