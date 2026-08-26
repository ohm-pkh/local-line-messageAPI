package dto

import "github.com/google/uuid"

type SecretResBody struct {
	AccessToken   uuid.UUID `json:"accessToken"`
	ChannelSecret uuid.UUID `json:"channelSecret"`
}
