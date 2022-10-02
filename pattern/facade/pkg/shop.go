package pkg

import (
	"errors"
	"time"
)

type Shop struct {
	Name     string
	Products []Product
}

func (s *Shop) Sell(user User, product string) error {
	println("[Магазин] Запрос к пользователю для получения остатка по карте")
	time.Sleep(time.Millisecond * 300)
	if err := user.CheckBalance(); err != nil {
		return err
	}

	println("[Магазин] Проверка возможности покупки товара покупателем", user.Name)
	time.Sleep(time.Millisecond * 300)

	for _, p := range s.Products {
		if product != p.Name {
			continue
		}

		if p.Price > user.GetBalance() {
			return errors.New("[Магазин] Недостаточно средств для покупки товара")
		}

		println("Товар", product, "куплен!")
	}

	return nil
}