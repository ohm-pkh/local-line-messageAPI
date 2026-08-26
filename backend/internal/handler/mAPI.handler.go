package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/ohm-pkh/local-line-messageAPI/internal/model/dto"
	apperror "github.com/ohm-pkh/local-line-messageAPI/internal/utils/app-error"
)

func (h *Handler) PushMessageHandler(c fiber.Ctx) error {
	auth := c.Get("Authorization")

	if auth == "" {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"missing authorization header",
		)
	}

	if !strings.HasPrefix(auth, "Bearer ") {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"invalid authorization header",
		)
	}

	token := strings.TrimPrefix(auth, "Bearer ")

	var req dto.LinePushMessage
	var appErr *apperror.AppError

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ResponseBody{
			Message: "Invalid request format.",
		})
	}

	if err := h.Service.ValidateAccessToken(token); err != nil {
		if errors.As(err, &appErr) {
			return c.Status(appErr.Code).JSON(fiber.Map{
				"error": appErr.Message,
			})
		}

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	res, err := h.Service.PushMessage(req)

	if err != nil {
		if errors.As(err, &appErr) {
			return c.Status(appErr.Code).JSON(fiber.Map{
				"error": appErr.Message,
			})
		}

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).JSON(res)

}
