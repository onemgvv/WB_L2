package pkg

import "fmt"

type plain struct {
	model       string
	engines     int
	luggageComp int
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
