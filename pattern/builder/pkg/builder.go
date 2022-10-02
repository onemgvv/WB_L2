package pkg

type OrderBuilder interface {
	SetCustomer(c Customer) OrderBuilder
	SetAddress(a Address) OrderBuilder
	SetProduct(p Product) OrderBuilder

	Build() *Order
}

type Builder struct {
	customer Customer
	address  Address
	product  Product
}

func (b *Builder) SetCustomer(c Customer) OrderBuilder {
	b.customer = c
	return b
}

func (b *Builder) SetAddress(a Address) OrderBuilder {
	b.address = a
	return b
}

func (b *Builder) SetProduct(p Product) OrderBuilder {
	b.product = p
	return b
}

func (b *Builder) Build() *Order {
	return NewOrder(b.customer, b.address, b.product)
}
