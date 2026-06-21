package splitwise

type Split interface {
	CalculateSplit(total int, users []*User) map[string]int
}

