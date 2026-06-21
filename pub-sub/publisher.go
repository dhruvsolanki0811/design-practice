package pubsub

import "github.com/google/uuid"

type Publisher struct {
	Id string
}

func NewPublisher() *Publisher {
	return &Publisher{
		Id: uuid.NewString(),
	}
}

func (p *Publisher) Publish(topic *Topic, message Message) {
	topic.OnMessage(message)
}
