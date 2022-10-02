package marketplace

type Warehouse struct {
	addressId   int
	packingId   int
	readyToShip bool
}

func NewWarehouse(addressId, packingId int) *Warehouse {
	return &Warehouse{addressId, packingId, false}
}

func (w *Warehouse) ChangeStatus() {
	w.readyToShip = true
}
