package service

import (
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/molpadia/molpastream/configs"
	"github.com/molpadia/molpastream/internal/domain/entity"
	"github.com/molpadia/molpastream/internal/domain/repository"
	"github.com/molpadia/molpastream/internal/infrastructure"
)

type BotService struct {
	adapter repository.BotAdapter
	repo    repository.BotRepository
}

// Create a new application service instance.
func NewBotService(repo repository.BotRepository) *BotService {
	adapter := infrastructure.NewBotAdapter(configs.AppConfig.LineChannelSecret, configs.AppConfig.LineAccessToken)
	return &BotService{adapter, repo}
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

// Send a push message to the LINE bot.
func (s *BotService) Send(userId, text string) error {
	msgs := []*entity.Message{{UserId: userId, Text: text, CreatedAt: time.Now()}}
	if err := s.adapter.BulkSend(userId, msgs); err != nil {
		return err
	}
	if err := s.repo.BulkSave(msgs); err != nil {
		return err
	}
	return nil
}
