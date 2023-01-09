package app

import "time"

type Message struct {
	UserId    string    `json:"userId" binding:"required"`
	Text      string    `json:"text" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
}
