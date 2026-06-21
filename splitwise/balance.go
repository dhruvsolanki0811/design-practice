package splitwise

type Balance struct {
	Amount     int
	ToUser     *User
	FromUser   *User
	Settlement []*Settlement
}

func NewBalance(amount int, toUser *User, fromUser *User) *Balance {
	return &Balance{
		Amount:   amount,
		ToUser:   toUser,
		FromUser: fromUser,
	}
}
