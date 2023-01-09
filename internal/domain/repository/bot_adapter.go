package repository

import "github.com/molpadia/molpastream/internal/domain/entity"

type BotAdapter interface {
	// Bulk send a list of push messages to LINE bot service.
	BulkSend(to string, msgs []*entity.Message) error
}
