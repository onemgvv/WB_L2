package pkg

type User struct {
	Name string
	*Card
}

func (u *User) GetBalance() float64 {
	return u.Balance
}