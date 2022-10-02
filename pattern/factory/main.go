package factory

import (
	"errors"
	"fmt"
	"log"
)

type Transport interface {
	Delivery()
}

type truck struct {
	trunk  int
	wheels int
}

type ship struct {
	hold     int
	boatType string
}

type plain struct {
	model       string
	engines     int
	luggageComp int
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

func (s *ship) Delivery() {
	fmt.Println("Ship is carrying on the ground")
}

func NewShip() Transport {
	return &ship{
		hold:     1000,
		boatType: "tanker",
	}
}

func (p *plain) Delivery() {
	fmt.Println("Plain is carrying in the air")
}

func NewPlain() Transport {
	return &plain{
		engines:     4,
		luggageComp: 700,
		model:       "Boeing Beluga",
	}
}

func getTransport(_type string) (Transport, error) {
	if _type == "ground" {
		return NewTruck(), nil
	} else if _type == "air" {
		return NewPlain(), nil
	} else if _type == "water" {
		return NewShip(), nil
	} else {
		return nil, errors.New("unknown type")
	}
}

func Factory() {
	deliverTransport, err := getTransport("ground")
	if err != nil {
		log.Fatalf("Delivery Transport error: %s", err.Error())
	}

	deliverTransport.Delivery()
}
