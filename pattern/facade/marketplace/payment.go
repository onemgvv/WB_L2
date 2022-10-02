package marketplace

type Payment struct {
	orderId    int
	totalPrice int
	isPayed    bool
}

func NewPayment(orderId, totalPrice int) *Payment {
	return &Payment{orderId, totalPrice, false}
}

func (p *Payment) ChangeStatus() {
	p.isPayed = true
}
