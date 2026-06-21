package pubsub

import "fmt"

type PrintSubsrciber struct {
	Id string
}

func (s *PrintSubsrciber) Handler(topic *Topic, message Message) {
	fmt.Println(message)
}

func (s *PrintSubsrciber) GetId() string {
	return s.Id
}
