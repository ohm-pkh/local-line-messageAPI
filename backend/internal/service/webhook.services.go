package service

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/ohm-pkh/local-line-messageAPI/internal/model/dto"
	apperror "github.com/ohm-pkh/local-line-messageAPI/internal/utils/app-error"
)

func (s *Service) RegisWebhook(url *dto.WebhookPath) error {
	s.webhook = url
	return nil
}

func (s *Service) RegisteredWebhook() (string, error) {
	if s.webhook == nil {
		return "", &apperror.AppError{
			Code:    fiber.StatusNotFound,
			Message: "Not Found Registered webhook.",
		}
	}

	return s.webhook.Path, nil
}

func (s *Service) ConstructWebhookEntity(event dto.ReqWebSocket) (*dto.LineWebhook, error) {
	userId := s.userId.String()
	quoteToken := uuid.New().String()

	webhookEntity := dto.LineWebhook{
		Destination: "Test_destination",
		Event: []dto.LineWebhookEvent{
			{
				Type:      "message",
				Mode:      "active",
				Timestamp: event.Timestamp,
				Source: &dto.LineWebhookSource{
					Type:   "user",
					UserId: &userId,
				},
				WebhookEventID: uuid.New().String(),
				DeliveryContext: dto.LineWebhookDeliveryContext{
					IsRedelivery: false,
				},
				ReplyToken: uuid.New().String(),
				Message: &dto.LineWebhookMessage{
					Id:              event.Data.MessageId,
					Type:            event.Type,
					QuoteToken:      &quoteToken,
					QuotedMessageId: event.QuoteMessageId,
					Text:            event.Data.Message,
				},
			},
		},
	}

	return &webhookEntity, nil
}

func (s *Service) TriggerWebhook(event *dto.LineWebhook) error {
	body, err := json.Marshal(event)
	if err != nil {
		return &apperror.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	if s.webhook == nil {
		return &apperror.AppError{Code: 200, Message: "No webhook registered."}
	}

	req, err := http.NewRequest(
		http.MethodPost,
		s.webhook.Path,
		bytes.NewBuffer(body),
	)
	if err != nil {
		return &apperror.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	mac := hmac.New(sha256.New, []byte(s.channelSecret.String()))
	mac.Write(body)

	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Line-Signature", signature)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &apperror.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return &apperror.AppError{
			Code:    resp.StatusCode,
			Message: fmt.Sprintf("Webhook returned status %d", resp.StatusCode),
		}
	}

	return nil
}
