package pkg

type Product struct {
	name  string
	count int
}

func NewProduct(name string, count int) *Product {
	return &Product{name, count}
}
