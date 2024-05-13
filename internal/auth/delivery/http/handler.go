package http

import (
	"car-api/internal/auth"
	authModel "car-api/internal/auth/model"
	"car-api/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	uc     auth.UC
	logger logger.ApiLogger
}

func New(uc auth.UC, logger logger.ApiLogger) auth.Handler {
	return &handler{
		uc:     uc,
		logger: logger,
	}
}

func (h *handler) SignIn() fiber.Handler {
	return func(c *fiber.Ctx) error {
		pass := c.FormValue("password")
		email := c.FormValue("email")

		_, err := h.uc.SignIn(&authModel.SignInParams{Email: email, Password: pass})
		if err != nil {
			return c.Status(fiber.StatusBadRequest).Render("sign_in", fiber.Map{})
		}
		c.Redirect("/", fiber.StatusFound)

		return c.Render("sign_in", fiber.Map{})
	}
}

func (h *handler) SignUp() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

func (h *handler) IndexView(c *fiber.Ctx) error {
	// Получаем сессию из контекста

	// Рендеринг шаблона index.tmpl с передачей значения isAuthenticated
	return c.Render("index", fiber.Map{
		"isAuthenticated": true,
		"ClientId":        c.Locals("user_id").(int64),
	})
}
