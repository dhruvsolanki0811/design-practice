package splitwise

type EqualSplit struct{}

func NewEqualSplit() *EqualSplit {
	return &EqualSplit{}
}

func (e *EqualSplit) CalculateSplit(amount int, participants []*User) map[string]int {
	share := make(map[string]int)
	for _, u := range participants {
		share[u.UserId] = amount / len(participants)
	}
	return share
}
