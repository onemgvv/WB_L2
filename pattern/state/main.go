package state

import (
	"log"
	"patterns/state/pkg"
)

func State() {
	vm := pkg.NewVendingMachine(1, 10)
	if err := vm.RequestItem(); err != nil {
		log.Fatal(err.Error())
	}

	if err := vm.InsertMoney(10); err != nil {
		log.Fatal(err.Error())
	}

	if err := vm.DispenseItem(); err != nil {
		log.Fatal(err.Error())
	}

	if err := vm.AddItem(2); err != nil {
		log.Fatal(err.Error())
	}

	println("")

	if err := vm.RequestItem(); err != nil {
		log.Fatal(err.Error())
	}

	if err := vm.InsertMoney(10); err != nil {
		log.Fatal(err.Error())
	}

	if err := vm.DispenseItem(); err != nil {
		log.Fatal(err.Error())
	}

	if err := vm.AddItem(2); err != nil {
		log.Fatal(err.Error())
	}
}
