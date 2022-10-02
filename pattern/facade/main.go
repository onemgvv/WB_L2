package facade

import "patterns/facade/pkg"

var (
	bank = pkg.Bank{
		Name:  "Закрытие",
		Cards: []pkg.Card{},
	}

	card1 = pkg.Card{
		Name:    "CRD-1",
		Balance: 200,
		Bank:    &bank,
	}

	card2 = pkg.Card{
		Name:    "CRD-2",
		Balance: 50,
		Bank:    &bank,
	}

	buyer = pkg.User{
		Name: "Buyer-1",
		Card: &card1,
	}

	buyer2 = pkg.User{
		Name: "Buyer-2",
		Card: &card2,
	}

	prod = pkg.Product{
		Name:  "laptop",
		Price: 125,
	}

	shop = pkg.Shop{
		Name:     "Re:Store",
		Products: []pkg.Product{prod},
	}
)

func Facade() {
	println("[Банк] Выпуск карт")
	bank.Cards = append(bank.Cards, card1, card2)

	println("[" + buyer.Name + "]")
	if err := shop.Sell(buyer, prod.Name); err != nil {
		println(err.Error())
	}

	println()

	println("[" + buyer.Name + "]")
	if err := shop.Sell(buyer2, prod.Name); err != nil {
		println(err.Error())
	}
}
