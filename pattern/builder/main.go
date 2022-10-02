package builder

import (
	"fmt"
	"patterns/builder/pkg"
)

func NewOrderBuilder() *pkg.Builder {
	return &pkg.Builder{}
}

func Builder() {
	customer := pkg.NewCustomer("Alex", "alex_customer@mail.ru", "+7 (999) 165-11-23")
	address := pkg.NewAddress("Moscow", "Arbat street", 12, 120123)
	product := pkg.NewProduct("Playstation 5", 1)

	oBuilder := NewOrderBuilder()

	order := oBuilder.SetCustomer(*customer).SetAddress(*address).SetProduct(*product).Build()
	fmt.Println("[ORDER]: ", order)
}
