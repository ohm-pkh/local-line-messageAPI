package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/ohm-pkh/local-line-messageAPI/internal/handler"
)

func Router(app *fiber.App, h handler.HandlerInt) {
	api := app.Group("api/v1")
	ws := app.Group("ws/v1")
	api.Get("/healthz", func(c fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	webhookRoute(api, h)
	secretRoute(api, h)
	webSoc(ws, h)
	MessageAPI(api, h)
}
