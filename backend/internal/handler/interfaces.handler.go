package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/ohm-pkh/local-line-messageAPI/internal/service"
	ws "github.com/ohm-pkh/local-line-messageAPI/internal/websocket"
)

type Handler struct {
	Service service.ServiceInt
	ws      *ws.Connection
}

type HandlerInt interface {
	HandleWebhookRegister(c fiber.Ctx) error
	GetRegisteredWebhook(c fiber.Ctx) error
	GetSecrets(c fiber.Ctx) error
	WebSocket() fiber.Handler
	TestWebSocket(c fiber.Ctx) error
	PushMessageHandler(c fiber.Ctx) error
}

func New(s service.ServiceInt, ws *ws.Connection) *Handler {
	return &Handler{
		Service: s,
		ws:      ws,
	}
}
