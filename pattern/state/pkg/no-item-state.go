package pkg

import "fmt"

type NoItemState struct {
	vm *VendingMachine
}

func (i *NoItemState) AddItem(count int) error {
	i.vm.incrementItemCount(count)
	i.vm.setState(i.vm.hasItem)
	return nil
}

func (i *NoItemState) RequestItem() error {
	return fmt.Errorf("Item out of stack")
}

func (i *NoItemState) InsertMoney(money int) error {
	return fmt.Errorf("Item out of stack")
}

func (i *NoItemState) DispenseItem() error {
	return fmt.Errorf("Item out of stack")
}