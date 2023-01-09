package infrastructure

import (
	"context"
	"log"

	"github.com/molpadia/molpastream/configs"
	"github.com/molpadia/molpastream/internal/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MsgCollection = "messages"

type BotRepository struct {
	ctx    context.Context
	client *mongo.Client
}

// Create a new repository instance.
func NewBotRepository(ctx context.Context, opts *options.ClientOptions) *BotRepository {
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}
	return &BotRepository{ctx, client}
}

// Bulk inserting a list of push messages to the persistence.
func (r *BotRepository) BulkSave(msgs []*entity.Message) error {
	var events []interface{}
	for _, msg := range msgs {
		events = append(events, msg)
	}
	if _, err := r.client.Database(configs.AppConfig.MongoDB).Collection(MsgCollection).InsertMany(r.ctx, events); err != nil {
		return err
	}
	return nil
}
