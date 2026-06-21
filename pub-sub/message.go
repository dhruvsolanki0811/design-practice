package pubsub

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id        string
	TimeStamp time.Time
	Content   string
}

func NewMessage(content string) Message {
	return Message{
		Id:        uuid.NewString(),
		TimeStamp: time.Now(),
		Content:   content,
	}
}
