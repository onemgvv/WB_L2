package marketplace

var id = 0

type Packing struct {
	Id        int
	orderId   int
	productId int
	count     int
}

func NewPacking(orderId, productId, count int) *Packing {
	id++
	return &Packing{id, orderId, productId, count}
}
