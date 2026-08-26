package service

import (
	"errors"
	"net/http"

	"github.com/ohm-pkh/local-line-messageAPI/internal/model/dto"
	apperror "github.com/ohm-pkh/local-line-messageAPI/internal/utils/app-error"
)

func (s *Service) TestSendMessage(msg string) error {
	sendData := map[string]any{
		"type": "text",
		"msg":  msg,
	}

	if err := s.SendToWebSocket(sendData); err != nil {
		return &apperror.AppError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

func (s *Service) SendToWebSocket(msg map[string]any) error {
	return s.ws.SendJSON(msg)
}

func (s *Service) ProcessEvent(event dto.ReqWebSocket) error {

	var appErr *apperror.AppError

	webhookEntity, err := s.ConstructWebhookEntity(event)
	if err != nil {
		return err
	}

	//send to db
	if webhookEntity == nil {
		return &apperror.AppError{
			Code:    http.StatusBadRequest,
			Message: "Entity Not Found.",
		}
	}
	err = s.repo.InsertMessage(webhookEntity.Event)
	if err != nil {
		if errors.As(err, &appErr) {
			if appErr.Code != 200 {
				return err
			}
		} else {
			return err
		}
	}

	//send to webhook
	err = s.TriggerWebhook(webhookEntity)

	if err != nil {
		if errors.As(err, &appErr) {
			if appErr.Code != 200 {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

func (s *Service) EventToWebsocketMessage(event []dto.LineWebhookEvent) ([]dto.ReqWebSocket, error) {
	var result []dto.ReqWebSocket

	for _, e := range event {
		temp := dto.ReqWebSocket{
			Type:           "text",
			QuoteMessageId: e.Message.QuotedMessageId,
			Timestamp:      e.Timestamp,
			Data: dto.ReqWebSocketData{
				Message:   e.Message.Text,
				MessageId: e.Message.Id,
			},
		}

		result = append(result, temp)
	}

	return result, nil
}
