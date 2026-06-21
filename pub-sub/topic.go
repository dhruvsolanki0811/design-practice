package pubsub

import (
	"sync"

	"github.com/google/uuid"
)

type Topic struct {
	Id          string
	Name        string
	Subsrcibers []Subsrciber
	mu          sync.RWMutex
}

func NewTopic(name string) *Topic {
	return &Topic{
		Id:          uuid.NewString(),
		Name:        name,
		Subsrcibers: []Subsrciber{},
	}
}

func (t *Topic) Subsrcibe(subscriber Subsrciber) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.Subsrcibers = append(t.Subsrcibers, subscriber)
}

func (t *Topic) Unsubsrcibe(subscriber Subsrciber) {
	t.mu.Lock()
	defer t.mu.Unlock()
	for i, sub := range t.Subsrcibers {
		if sub.GetId() == subscriber.GetId() {
			t.Subsrcibers = append(t.Subsrcibers[:i], t.Subsrcibers[i+1:]...)
			return
		}
	}
}

func (t *Topic) OnMessage(message Message) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	for _, sub := range t.Subsrcibers {
		sub.Handler(t, message)
	}
}
