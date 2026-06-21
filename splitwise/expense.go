package splitwise

import "github.com/google/uuid"

type Expense struct {
	Id           string
	Amount       int
	Participants []*User
	Group        *Group
	Split        Split
	PaidBy       *User
	Share        map[string]int
}

func NewExpense(amount int, participants []*User, split Split, paidBy *User, group *Group) *Expense {
	share := split.CalculateSplit(amount, participants)

	return &Expense{
		Id:           uuid.NewString(),
		Amount:       amount,
		Participants: append(participants, paidBy),
		Group:        group,
		Split:        split,
		PaidBy:       paidBy,
		Share:        share,
	}
}
