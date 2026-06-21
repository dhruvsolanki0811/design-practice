package pubsub

type Subsrciber interface {
	Handler(topic *Topic, message Message)
	GetId() string
}
