package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/fiber/v3"
	"github.com/ohm-pkh/local-line-messageAPI/internal/model/dto"
	apperror "github.com/ohm-pkh/local-line-messageAPI/internal/utils/app-error"
)

func (h *Handler) WebSocket() fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		h.ws.Set(c)

		defer h.ws.Clear()

		for {
			_, p, err := c.ReadMessage()
			if err != nil {
				return
			}
			var req dto.ReqWebSocket

			if err := json.Unmarshal(p, &req); err != nil {
				fmt.Println("Invalid message:", err)
				continue
			}

			fmt.Printf("Type: %s\n", req.Type)

			if req.QuoteMessageId != nil {
				fmt.Println("QuoteMessageId:", *req.Data.Message)
			}

			if req.Data.Message != nil {
				fmt.Println("Message:", *req.Data.Message)
			}

			if req.Data.Stage != nil {
				fmt.Println("Stage:", *req.Data.Stage)
			}

			if req.Data.Status != nil {
				fmt.Println("Status:", *req.Data.Status)
			}

			err = h.Service.ProcessEvent(req)

			if err != nil {
				errMsg := err.Error()
				h.ws.SendJSON(dto.ReqWebSocket{
					Type:      "Error",
					Timestamp: time.Now().UnixMicro(),
					Data: dto.ReqWebSocketData{
						Message: &errMsg,
					},
				})
			}
		}
	})
}

func (h *Handler) TestWebSocket(c fiber.Ctx) error {
	var req dto.ReqWebSocket
	var appErr *apperror.AppError

	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseBody{
			Message: "Invalid request format.",
		})
	}

	if req.Data.Message == nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if err := h.Service.TestSendMessage(*req.Data.Message); err != nil {
		if errors.As(err, &appErr) {
			return c.Status(appErr.Code).JSON(fiber.Map{
				"error": appErr.Message,
			})
		}

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusAccepted)
}
