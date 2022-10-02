package pkg

import "fmt"

type HasMoneyState struct {
	vm *VendingMachine
}

func (i *HasMoneyState) AddItem(count int) error {
	return fmt.Errorf("Item dispence in process")
}

func (i *HasMoneyState) RequestItem() error {
	return fmt.Errorf("Item dispence in process")
}

func (i *HasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("Item dispence in process")
}

func (i *HasMoneyState) DispenseItem() error {
	fmt.Println("Dispancing item")
	i.vm.decrementItemCount(1)
	if i.vm.itemCount == 0 {
		i.vm.setState(i.vm.noItem)
	} else {
		i.vm.setState(i.vm.hasItem)
	}

	return nil
}