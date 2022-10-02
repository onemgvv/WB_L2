package builder

import "fmt"

type customer struct {
	name  string
	phone string
	email string
}

type address struct {
	city   string
	street string
	home   int
	zip    int
}

type product struct {
	name  string
	count int
}

type order struct {
	customer
	address
	product
}

type OrderBuilder interface {
	SetCustomer(c customer) OrderBuilder
	SetAddress(a address) OrderBuilder
	SetProduct(p product) OrderBuilder

	Build() order
}

type orderBuilder struct {
	customer
	address
	product
}

func (b *orderBuilder) SetCustomer(c customer) OrderBuilder {
	b.customer = c
	return b
}

func (b *orderBuilder) SetAddress(a address) OrderBuilder {
	b.address = a
	return b
}

func (b *orderBuilder) SetProduct(p product) OrderBuilder {
	b.product = p
	return b
}

func (b *orderBuilder) Build() order {
	return order{
		customer: b.customer,
		address:  b.address,
		product:  b.product,
	}
}

func NewOrderBuilder() *orderBuilder {
	return &orderBuilder{}
}

func Builder() {
	customer := customer{
		name:  "Alex",
		email: "alex_customer@mail.ru",
		phone: "+7 (999) 165-11-23",
	}

	address := address{
		city:   "Moscow",
		street: "Arbat street",
		home:   12,
		zip:    100100,
	}

	product := product{
		name:  "Playstation 5",
		count: 1,
	}

	oBuilder := NewOrderBuilder()

	order := oBuilder.SetCustomer(customer).SetAddress(address).SetProduct(product).Build()
	fmt.Println("[ORDER]: ", order)
}
