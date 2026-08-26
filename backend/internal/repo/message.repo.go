package repo

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/ohm-pkh/local-line-messageAPI/internal/model/dto"
	apperror "github.com/ohm-pkh/local-line-messageAPI/internal/utils/app-error"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (r *Repo) InsertMessage(message []dto.LineWebhookEvent) error {
	ctx := context.Background()

	docs := make([]any, len(message))
	for i, event := range message {
		docs[i] = event
	}

	_, err := r.db.Collection("messages").InsertMany(ctx, docs)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) MapQuoteToken(token *string) (*string, error) {
	fmt.Println("token:", *token)

	var message dto.LineWebhookEvent

	err := r.db.Collection("messages").FindOne(
		context.Background(),
		bson.M{
			"message.quoteToken": *token,
		},
	).Decode(&message)

	if err != nil {
		fmt.Println("mongo error:", err)

		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, &apperror.AppError{
				Code:    http.StatusNotFound,
				Message: "message not found",
			}
		}

		return nil, err
	}

	fmt.Println("found message:", message.Message.Id)

	return message.Message.Id, nil
}
