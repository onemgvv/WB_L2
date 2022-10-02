package pkg

type Customer struct {
	name  string
	phone string
	email string
}

func NewCustomer(name, phone, email string) *Customer {
	return &Customer{name, phone, email}
}
