package handler

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/ohm-pkh/local-line-messageAPI/internal/model/dto"
	apperror "github.com/ohm-pkh/local-line-messageAPI/internal/utils/app-error"
)

func (h *Handler) HandleWebhookRegister(c fiber.Ctx) error {
	var req dto.WebhookPath

	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseBody{
			Message: "Invalid request format.",
		})
	}

	log.Printf("REQ: %+v", req)
	log.Printf("RAW: %s", c.Body())

	if err := h.Service.RegisWebhook(req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ResponseBody{
			Message: "Fail register webhook",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(dto.ResponseBody{
		Message: "Register Successful.",
	})
}

func (h *Handler) GetRegisteredWebhook(c fiber.Ctx) error {
	var appErr *apperror.AppError

	url, err := h.Service.RegisteredWebhook()
	if err != nil {
		if errors.As(err, &appErr) {
			return c.Status(appErr.Code).JSON(fiber.Map{
				"error": appErr.Message,
			})
		}

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(dto.ResponseBody{
		Message: "webhook is: " + url,
	})
}
