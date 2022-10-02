package pkg

import "fmt"

type truck struct {
	trunk  int
	wheels int
}

func (t *truck) Delivery() {
	fmt.Println("Truck is carrying on ground")
}

func NewTruck() Transport {
	return &truck{
		trunk:  100,
		wheels: 10,
	}
}
