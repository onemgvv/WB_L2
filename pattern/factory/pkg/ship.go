package pkg

import "fmt"

type ship struct {
	hold     int
	boatType string
}

func (s *ship) Delivery() {
	fmt.Println("Ship is carrying on the ground")
}

func NewShip() Transport {
	return &ship{
		hold:     1000,
		boatType: "tanker",
	}
}
