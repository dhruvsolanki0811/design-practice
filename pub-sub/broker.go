package pubsub

import "sync"

type Broker struct {
	Publishers []*Publisher
	Topics     []*Topic
	mu         sync.Mutex
}

var (
	broker *Broker
	once   sync.Once
)

func NewBroker() *Broker {
	once.Do(func() {
		broker = &Broker{
			Publishers: []*Publisher{},
			Topics:     []*Topic{},
		}
	})
	return broker
}

func (b *Broker) CreatePublisher() *Publisher {
	b.mu.Lock()
	defer b.mu.Unlock()
	publisher := NewPublisher()
	b.Publishers = append(b.Publishers, publisher)
	return publisher
}

func (b *Broker) CreateTopic(name string) *Topic {
	b.mu.Lock()
	defer b.mu.Unlock()
	topic := NewTopic(name)
	b.Topics = append(b.Topics, topic)
	return topic
}

func (b *Broker) SubsrcibeTopic(sub Subsrciber, topic *Topic) {
	b.mu.Lock()
	defer b.mu.Unlock()
	for _, t := range b.Topics {
		if topic.Id == t.Id {
			t.Subsrcibe(sub)
		}
	}
}

func (b *Broker) PublishMessage(message Message, pub *Publisher, topic *Topic) {
	pub.Publish(topic, message)
}
