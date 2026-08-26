package repo

import (
	"github.com/ohm-pkh/local-line-messageAPI/internal/model/dto"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repo struct {
	db *mongo.Database
}

type RepoInt interface {
	InsertMessage(message []dto.LineWebhookEvent) error
	MapQuoteToken(token *string) (*string, error)
}

func New(db *mongo.Database) *Repo {
	return &Repo{
		db: db,
	}
}
