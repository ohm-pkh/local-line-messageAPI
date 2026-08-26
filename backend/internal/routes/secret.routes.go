package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/ohm-pkh/local-line-messageAPI/internal/handler"
)

func secretRoute(router fiber.Router, h handler.HandlerInt) {
	api := router.Group("secret")

	api.Get("/", h.GetSecrets)
}
