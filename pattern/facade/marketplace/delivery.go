package marketplace

import "time"

type Delivery struct {
	orderId      int
	address      string
	deliveryDate time.Time
	delivered    bool
}

func NewDelivery(orderId int, address string, delDate time.Time) *Delivery {
	return &Delivery{orderId, address, delDate, false}
}

func (d *Delivery) ChangeStatus() {
	d.delivered = true
}
