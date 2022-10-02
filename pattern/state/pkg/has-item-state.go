package pkg

import "fmt"

type HasItemState struct {
	vm *VendingMachine
}

func (i *HasItemState) AddItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vm.incrementItemCount(count)
	return nil
}

func (i *HasItemState) RequestItem() error {
	if i.vm.itemCount == 0 {
		i.vm.setState(i.vm.noItem)
		return fmt.Errorf("No item present")
	}

	fmt.Println("Item requested")
	i.vm.setState(i.vm.itemRequest)
	return nil
}

func (i *HasItemState) InsertMoney(money int) error {
	return fmt.Errorf("Please select item first")
}

func (i *HasItemState) DispenseItem() error {
	return fmt.Errorf("Please select item first")
}
