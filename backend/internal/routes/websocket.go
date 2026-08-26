package routes

import (
	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/fiber/v3"
	"github.com/ohm-pkh/local-line-messageAPI/internal/handler"
)

func webSoc(router fiber.Router, h handler.HandlerInt) {
	router.Use("/connect", func(c fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	router.Get("/connect", h.WebSocket())

	router.Post("/test", h.TestWebSocket)

}
