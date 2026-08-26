package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/ohm-pkh/local-line-messageAPI/internal/handler"
)

func MessageAPI(router fiber.Router, h handler.HandlerInt) {
	api := router.Group("bot/message")

	api.Post("/push", h.PushMessageHandler)
}
