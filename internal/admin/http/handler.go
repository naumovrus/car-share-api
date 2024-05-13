package http

import (
	"car-api/internal/admin"
	"car-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	mw middleware.MW
}

func New() admin.Handler {
	return &handler{}
}

func (r *handler) AdminView() fiber.Handler {
	return func(c *fiber.Ctx) error {

		return c.Render("index", fiber.Map{
			"IsAdmin":         nil,
			"ClientId":        nil,
			"IsAuthenticated": nil,
		})
	}
}
