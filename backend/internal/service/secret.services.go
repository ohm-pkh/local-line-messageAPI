package service

import (
	"net/http"

	"github.com/ohm-pkh/local-line-messageAPI/internal/model/dto"
	apperror "github.com/ohm-pkh/local-line-messageAPI/internal/utils/app-error"
)

func (s *Service) FindSecret() *dto.SecretResBody {
	return &dto.SecretResBody{
		AccessToken:   s.accessToken,
		ChannelSecret: s.channelSecret,
	}
}

func (s *Service) ValidateAccessToken(token string) error {
	if token == "" {
		return &apperror.AppError{
			Code:    http.StatusUnauthorized,
			Message: "missing access token",
		}
	}

	if token != s.accessToken.String() {
		return &apperror.AppError{
			Code:    http.StatusUnauthorized,
			Message: "invalid access token",
		}
	}
	return nil
}
