package pkg

import (
	"fmt"
)

type Transport interface {
	Delivery()
}

type DeliveryType string

const (
	Ground DeliveryType = "ground"
	Water               = "water"
	Air                 = "air"
)

func NewTransport(delType DeliveryType) Transport {
	if delType == "ground" {
		return NewTruck()
	} else if delType == "air" {
		return NewPlain()
	} else if delType == "water" {
		return NewShip()
	} else {
		fmt.Println("[Transport Factory] Unknown type", delType)
		return nil
	}
}
