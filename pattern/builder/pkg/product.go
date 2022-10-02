package pkg

type Order struct {
	customer Customer
	address  Address
	product  Product
}

func NewOrder(customer Customer, address Address, product Product) *Order {
	return &Order{customer, address, product}
}
