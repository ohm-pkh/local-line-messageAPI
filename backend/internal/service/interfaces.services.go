package service

import (
	"github.com/google/uuid"
	"github.com/ohm-pkh/local-line-messageAPI/internal/model/dto"
	"github.com/ohm-pkh/local-line-messageAPI/internal/repo"
	ws "github.com/ohm-pkh/local-line-messageAPI/internal/websocket"
)

type Service struct {
	webhook       *dto.WebhookPath
	accessToken   uuid.UUID
	channelSecret uuid.UUID
	userId        uuid.UUID
	repo          repo.RepoInt
	ws            *ws.Connection
}

type ServiceInt interface {
	RegisWebhook(url dto.WebhookPath) error
	RegisteredWebhook() (string, error)
	FindSecret() *dto.SecretResBody
	TestSendMessage(msg string) error
	SendToWebSocket(msg map[string]any) error
	ProcessEvent(event dto.ReqWebSocket) error
	ValidateAccessToken(token string) error
	PushMessage(req dto.LinePushMessage) (*dto.PushMessageResponse, error)
}

func New(accessToken uuid.UUID, channelSecret uuid.UUID, ws *ws.Connection, repo repo.RepoInt) *Service {
	return &Service{
		accessToken:   accessToken,
		channelSecret: channelSecret,
		userId:        uuid.New(),
		repo:          repo,
		ws:            ws,
	}
}
