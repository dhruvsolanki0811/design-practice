package splitwise

import (
	"errors"
	"strings"
	"sync"
)

type Splitwise struct {
	Users    map[string]*User
	Groups   map[string]*Group
	Balances map[string]map[string]*Balance
	mu       sync.RWMutex
}

var (
	splitwiseApp *Splitwise
	once         sync.Once
)

func NewSplitwiseApp() *Splitwise {
	once.Do(func() {
		splitwiseApp = &Splitwise{
			Users:    make(map[string]*User),
			Groups:   make(map[string]*Group),
			Balances: make(map[string]map[string]*Balance),
		}
	})
	return splitwiseApp
}

func (s *Splitwise) CreateNewUser(name string) *User {
	s.mu.Lock()
	defer s.mu.Unlock()
	user := NewUser(name)
	s.Users[user.UserId] = user
	return user
}

func (s *Splitwise) CreateNewGroup(name string, creator *User) *Group {
	s.mu.Lock()
	defer s.mu.Unlock()
	group := NewGroup(name, creator)
	s.Groups[group.Id] = group

	return group
}

func (s *Splitwise) AddMember(user *User, group *Group) {
	s.mu.Lock()
	defer s.mu.Unlock()
	group.AddMember(user)
}

func (s *Splitwise) AddExpenses(amount int, participants []*User, splitType string, paidBy *User, group *Group) *Expense {
	s.mu.Lock()
	defer s.mu.Unlock()
	split := s.CreateSplit(splitType)
	expense := NewExpense(amount, participants, split, paidBy, group)
	group.AddGroupExpense(expense)
	for _, participant := range participants {
		if participant.UserId == paidBy.UserId {
			continue
		}
		if s.Balances[paidBy.UserId] == nil {
			s.Balances[paidBy.UserId] = make(map[string]*Balance)
		}
		balance, ok := s.Balances[paidBy.UserId][participant.UserId]
		if !ok {
			balance = NewBalance(0, participant, paidBy)
			s.Balances[paidBy.UserId][participant.UserId] = balance
		}
		balance.Amount += expense.Share[participant.UserId]
	}
	return expense
}

func (s *Splitwise) CreateSplit(split string) Split {
	switch strings.ToLower(split) {
	case "equal":
		return NewEqualSplit()
	default:
		return NewEqualSplit()
	}
}

func (s *Splitwise) SettleUp(paidBy *User, paidTo *User, amount int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if amount <= 0 {
		return errors.New("error")
	}
	user, ok := s.Balances[paidTo.UserId]
	if !ok || user == nil {
		return errors.New("error")
	}
	balance, ok := user[paidBy.UserId]
	if !ok || balance == nil {
		return errors.New("error")
	}

	if balance.Amount < amount {
		return errors.New("error")
	}

	balance.Amount -= amount

	settlement := NewSettlement(amount)
	balance.Settlement = append(balance.Settlement, settlement)

	return nil
}

func (s *Splitwise) GetBalance(user *User) map[string]int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	balances := map[string]int{}
	users, _ := s.Balances[user.UserId]
	for userId, balance := range users {
		balances[userId] = balance.Amount
	}

	for creditorId, innerMap := range s.Balances {
		if creditorId == user.UserId {
			continue
		}
		if balance, ok := innerMap[user.UserId]; ok {
			balances[creditorId] -= balance.Amount
		}
	}
	return balances
}

func (s *Splitwise) GetBalanceWithUser(user1, user2 *User) int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	amount := 0
	innerMap, _ := s.Balances[user1.UserId]
	balance, ok := innerMap[user2.UserId]
	if ok {
		amount += balance.Amount
	}

	innerMap, _ = s.Balances[user2.UserId]
	balance, ok = innerMap[user1.UserId]
	if ok {
		amount -= balance.Amount
	}

	return amount
}

func (s *Splitwise) GetGroupExpenses(group *Group) ([]*Expense, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if group == nil {
		return nil, errors.New("group not found")
	}
	return group.Expenses, nil
}

func (s *Splitwise) GetUserHistory(user *User) []*Settlement {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var history []*Settlement
	for _, innerMap := range s.Balances {
		for _, balance := range innerMap {
			if balance.FromUser.UserId == user.UserId || balance.ToUser.UserId == user.UserId {
				history = append(history, balance.Settlement...)
			}
		}
	}
	return history
}
