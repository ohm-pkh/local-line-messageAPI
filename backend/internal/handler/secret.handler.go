package handler

import "github.com/gofiber/fiber/v3"

func (h *Handler) GetSecrets(c fiber.Ctx) error {

	res := h.Service.FindSecret()

	return c.Status(fiber.StatusOK).JSON(&res)
}
