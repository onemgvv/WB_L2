package pkg

import (
	"errors"
	"time"
)

type Bank struct {
	Name  string
	Cards []Card
}

func (b *Bank) CheckBalance(card string) error {
	println("[Банк] Получение остатка по карте", card)
	time.Sleep(time.Microsecond * 400)

	for _, v := range b.Cards {
		if v.Name != card {
			continue
		}

		if v.Balance <= 0 {
			return errors.New("[Банк] Недостаточно средств")
		}
	}
	println("[Банк] Остаток положительный")
	return nil
}