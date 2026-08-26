package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/ohm-pkh/local-line-messageAPI/internal/handler"
)

func webhookRoute(router fiber.Router, h handler.HandlerInt) {
	api := router.Group("webhook")

	api.Post("/register", h.HandleWebhookRegister)
	api.Get("/register", h.GetRegisteredWebhook)
}
