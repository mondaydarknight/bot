package service

import (
	"context"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/molpadia/molpastream/configs"
	"github.com/molpadia/molpastream/internal/domain/entity"
	"github.com/molpadia/molpastream/internal/domain/repository"
	"github.com/molpadia/molpastream/internal/infrastructure"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BotService struct {
	repo repository.BotRepository
}

// Create a new application service instance.
func NewBotService(ctx context.Context) *BotService {
	repo := infrastructure.NewBotRepository(ctx, options.Client().ApplyURI(configs.AppConfig.MongoUri))
	return &BotService{repo}
}

// Store a push message from the webhook.
func (s *BotService) Webhook(events []*linebot.Event) error {
	var msgs []*entity.Message
	for _, e := range events {
		if e.Type == linebot.EventTypeMessage {
			msgs = append(msgs, &entity.Message{
				UserId:    e.Source.UserID,
				Text:      e.Message.(*linebot.TextMessage).Text,
				CreatedAt: time.Now(),
			})
		}
	}
	return s.repo.BulkSave(msgs)
}
