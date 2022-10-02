package facade

import (
	"fmt"
	"patterns/facade/marketplace"
	"time"
)

var (
	orderId   = 0
	productId = 0
	addressId = 0
)

type orderFacade struct {
	marketplace.Account
	marketplace.Payment
	marketplace.Tax
	marketplace.Product
	marketplace.Packing
	marketplace.Warehouse
	marketplace.Delivery
}

func newOrderFacade(orderId, totalPrice, productId, addressId int, productName string) *orderFacade {
	acc := marketplace.NewAccount("name", 1901, 0)
	pay := marketplace.NewPayment(orderId, totalPrice)
	tx := marketplace.NewTax(orderId, int(float64(totalPrice)*0.15))
	prod := marketplace.NewProduct(productId, productName)
	pack := marketplace.NewPacking(orderId, productId, 1)
	wrh := marketplace.NewWarehouse(addressId, pack.Id)
	del := marketplace.NewDelivery(orderId, "Address street", time.Now().Add(time.Hour*5))

	return &orderFacade{
		Account:   *acc,
		Payment:   *pay,
		Tax:       *tx,
		Product:   *prod,
		Packing:   *pack,
		Warehouse: *wrh,
		Delivery:  *del,
	}
}

func (f *orderFacade) makeOrder() {
	orderId++
	productId++
	addressId++
	f.Authorize("name", 1901)
	f.Payment.ChangeStatus()
	f.Tax.MarkAsPayed()
	f.Warehouse.ChangeStatus()
	f.Delivery.ChangeStatus()
}

func Facade() {
	order := newOrderFacade(orderId, 500, productId, addressId, "phone case")
	order.makeOrder()
	fmt.Println()
}
