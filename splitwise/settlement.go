package splitwise

import (
	"time"

	"github.com/google/uuid"
)

type Settlement struct {
	Id        string
	Amount    int
	Timestamp time.Time
}

func NewSettlement(amount int) *Settlement {
	return &Settlement{
		Id:        uuid.NewString(),
		Amount:    amount,
		Timestamp: time.Now(),
	}
}
