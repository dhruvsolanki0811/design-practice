package splitwise

import (
	"github.com/google/uuid"
)

type Group struct {
	Id       string
	Name     string
	Users    []*User
	Expenses []*Expense
}

func NewGroup(name string, user *User) *Group {
	group := &Group{
		Id:    uuid.NewString(),
		Name:  name,
		Users: []*User{user},
	}
	return group
}

func (g *Group) AddMember(User *User) {
	g.Users = append(g.Users, User)
}

func (g *Group) AddGroupExpense(expense *Expense) {
	g.Expenses = append(g.Expenses, expense)
}
