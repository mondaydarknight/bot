package entity

import "time"

type Message struct {
	UserId    string `bson:"userId,omitempty"`
	Text      string `bson:"text,omitempty"`
	CreatedAt time.Time
}
