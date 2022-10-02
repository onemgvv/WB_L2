package pkg

import "fmt"

type ItemRequestState struct {
	vm *VendingMachine
}

func (i *ItemRequestState) AddItem(count int) error {
	return fmt.Errorf("Item dispence in process")
}

func (i *ItemRequestState) RequestItem() error {
		return fmt.Errorf("Item already requested")
}

func (i *ItemRequestState) InsertMoney(money int) error {
	if money < i.vm.itemPrice {
		return fmt.Errorf("Inserted money is less, please insert one more: %d", i.vm.itemPrice - money)
	}

	fmt.Println("Money inserted!")
	i.vm.setState(i.vm.hasMoney)
	return nil
}

func (i *ItemRequestState) DispenseItem() error {
	return fmt.Errorf("Please insert money for get item")
}