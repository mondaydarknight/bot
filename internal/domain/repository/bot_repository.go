package repository

import "github.com/molpadia/molpastream/internal/domain/entity"

type BotRepository interface {
	// Bulk inserting a list of push messages to the persistence.
	BulkSave(msgs []*entity.Message) error
}
