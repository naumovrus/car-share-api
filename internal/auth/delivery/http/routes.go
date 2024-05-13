package http

import (
	"car-api/internal/auth"
	"car-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func MapAuthRoutes(router fiber.Router, mw middleware.MW, h auth.Handler) {
	router.Post("sign_in", h.SignIn())
	// router.Post("sign_up", h.SignUp())
}