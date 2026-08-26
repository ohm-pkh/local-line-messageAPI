package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/ohm-pkh/local-line-messageAPI/internal/model/dto"
)

func (s *Service) PushMessage(req dto.LinePushMessage) (*dto.PushMessageResponse, error) {
	//map quoteToken to messageId
	for i := range req.Message {
		if req.Message[i].QuoteToken == nil {
			continue
		}
		messageID, err := s.repo.MapQuoteToken(req.Message[i].QuoteToken)
		if err != nil {
			return nil, err
		}

		req.Message[i].QuotedMessageId = messageID
		req.Message[i].QuoteToken = nil
	}

	event, err := s.ConstructLineEvent(req.Message)
	if err != nil {
		return nil, err
	}

	err = s.repo.InsertMessage(event)
	if err != nil {
		return nil, err
	}

	//construct response
	res, err := s.ConstructPushMessageRes(event)
	if err != nil {
		return nil, err
	}

	//construct ws
	wsReq, err := s.EventToWebsocketMessage(event)
	if err != nil {
		return nil, err
	}

	//push to ws
	if req.To != s.userId.String() {
		return &res, nil
	}
	for _, r := range wsReq {
		s.ws.SendJSON(r)
	}

	return &res, nil
}

func (s *Service) ConstructLineEvent(req []dto.LineWebhookMessage) ([]dto.LineWebhookEvent, error) {
	var result []dto.LineWebhookEvent

	for _, m := range req {
		messageID := uuid.NewString()
		quoteToken := uuid.NewString()

		m.Id = &messageID
		m.QuoteToken = &quoteToken

		temp := dto.LineWebhookEvent{
			Type:           "message",
			Mode:           "active",
			Timestamp:      time.Now().UnixMilli(),
			WebhookEventID: uuid.NewString(),

			DeliveryContext: dto.LineWebhookDeliveryContext{
				IsRedelivery: false,
			},

			ReplyToken: uuid.NewString(),
			Message:    &m,
		}

		result = append(result, temp)
	}

	return result, nil
}

func (s *Service) ConstructPushMessageRes(events []dto.LineWebhookEvent) (dto.PushMessageResponse, error) {
	var sentMessages []dto.SentMessage

	for _, e := range events {
		temp := dto.SentMessage{
			Id:         *e.Message.Id,
			QuoteToken: *e.Message.QuoteToken,
		}

		sentMessages = append(sentMessages, temp)
	}

	return dto.PushMessageResponse{
		SentMessages: sentMessages,
	}, nil
}
