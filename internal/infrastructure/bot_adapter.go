package infrastructure

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/molpadia/molpastream/internal/domain/entity"
)

type BotAdapter struct {
	bot *linebot.Client
}

// Create a new adapter instance.
func NewBotAdapter(secret, token string) *BotAdapter {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	bot, err := linebot.New(secret, token, linebot.WithHTTPClient(client))
	if err != nil {
		log.Fatal(err)
	}
	return &BotAdapter{bot}
}

// Bulk send a list of push messages to LINE bot service.
func (a *BotAdapter) BulkSend(to string, msgs []*entity.Message) error {
	var messages []linebot.SendingMessage
	for _, msg := range msgs {
		messages = append(messages, linebot.NewTextMessage(msg.Text))
	}
	if _, err := a.bot.PushMessage(to, messages...).Do(); err != nil {
		return err
	}
	return nil
}
