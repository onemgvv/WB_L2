package pkg

type VendingMachine struct {
	hasItem State
	itemRequest State
	hasMoney State
	noItem State
	currentState State
	itemCount int
	itemPrice int
}

func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}

	his := &HasItemState{v}
	irs := &ItemRequestState{v}
	hms := &HasMoneyState{v}
	nis := &NoItemState{v}

	v.setState(his)
	v.hasItem = his
	v.noItem = nis
	v.itemRequest = irs
	v.hasMoney = hms
	
	return v
}

func (vm *VendingMachine) AddItem(count int) error {
	return vm.currentState.AddItem(count)
}

func (vm *VendingMachine) RequestItem() error {
	return vm.currentState.RequestItem()
}

func (vm *VendingMachine) InsertMoney(money int) error {
	return vm.currentState.InsertMoney(money)
}

func (vm *VendingMachine) DispenseItem() error {
	return vm.currentState.DispenseItem()
}

func (vm *VendingMachine) setState(state State) {
	vm.currentState = state
}

func (vm *VendingMachine) incrementItemCount(count int) {
	vm.itemCount += count
}
func (vm *VendingMachine) decrementItemCount(count int) {
	vm.itemCount -= count
}