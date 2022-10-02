package marketplace

type Tax struct {
	orderId   int
	taxAmount int
	isPayed   bool
}

func NewTax(orderId, taxAmount int) *Tax {
	return &Tax{orderId, taxAmount, false}
}

func (t *Tax) MarkAsPayed() {
	t.isPayed = true
}
