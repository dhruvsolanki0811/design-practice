package splitwise

import "github.com/google/uuid"

type User struct {
	UserId string
	Name   string
}

func NewUser(Name string) *User {
	return &User{
		UserId: uuid.New().String(),
		Name:   Name,
	}
}
