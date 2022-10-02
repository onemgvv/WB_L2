package pkg

import "time"

type Card struct {
	Name string
	Balance float64
	*Bank
}

func (c *Card) CheckBalance() error {
	println("[Карта] Запрос в банк для проверки остатка")
	time.Sleep(time.Millisecond * 800)

	return c.Bank.CheckBalance(c.Name)
}