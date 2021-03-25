package card

import (
	"bank/pkg/bank/types"
)

const (
	withdrawLimit types.Money = 20_000_00
	depositLimit types.Money = 50_000_00
	bonusLimit types.Money = 5_000_00
)

func IssueCard(currency types.Currency, color string, name string) types.Card {
	newCard := types.Card{
		ID:         1000,
		PAN:        "5058 xxxx xxxx 0001",
		Balance:    0,
		MinBalance: 0,
		Currency:   currency,
		Color:      color,
		Name:       name,
		Active:     true,
	}

	return newCard
}

func Withdraw(card *types.Card, amount types.Money) {
	if amount < 0 {
		return
	}
	if amount > withdrawLimit {
		return
	}
	if !(*card).Active {
		return
	}
	if (*card).Balance < amount {
		return
	}

	(*card).Balance -= amount
}

func Deposit(card *types.Card, amount types.Money) {
	if !card.Active {
		return
	}
	if amount > depositLimit {
		return
	}
	if amount < 0 {
		return
	}

	card.Balance += amount
}

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if !card.Active {
		return
	}
	if card.Balance <= 0 {
		return
	}
	
	bonus := card.MinBalance / 100 * types.Money(percent) * types.Money(daysInMonth) / types.Money(daysInYear)

	if bonus > bonusLimit {
		return
	}

	card.Balance += bonus
}

// Total вычисляет общую сумму средств на всех картах.
// Отрицательные суммы и суммы на заблокированных картах игнорируются.
func Total(cards []types.Card) types.Money {
	total := types.Money(0)
	for _, card := range cards {
		if !card.Active {
			continue
		}
		
		if card.Balance <= 0 {
			continue
		}
		
		total += card.Balance
	}
	
	return total
}