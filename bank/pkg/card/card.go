package card

import "bank/pkg/bank/types"

func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}

	return card
}

func Withdraw(card types.Card, amount types.Money) types.Card {

	if !card.Active {
		return card
	}

	if amount <= 0 {
		return card
	}

	if amount > card.Balance {
		return card
	}

	if amount > 20_000_00 {
		return card
	}

	card.Balance -= amount

	return card
}

func Deposit(card *types.Card, amount types.Money) {

	if !card.Active {
		return
	}

	if amount > 50_000_00 {
		return
	}

	card.Balance += amount
}

func Total(cards []types.Card) types.Money {
	var sum types.Money

	for _, card := range cards {
		if !card.Active {
			continue
		}

		if card.Balance < 0 {
			continue
		}

		sum += card.Balance
	}

	return sum
}