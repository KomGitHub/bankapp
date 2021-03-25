package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

/*func ExampleWithdraw_positive() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)

	// Output: 1000000
}

func ExampleWithdraw_noMoney() {
	card := types.Card{Balance: 0, Active: true}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)

	// Output: 0
}

func ExampleWithdraw_inactive() {
	card := types.Card{Balance: 20_000_00, Active: false}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)

	// Output: 2000000
}

func ExampleWithdraw_limit() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Withdraw(&card, 20_000_01)
	fmt.Println(card.Balance)

	// Output: 2000000
}

func ExampleDeposit_normal() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&card, 10_000_00)
	fmt.Println(card.Balance)

	// Output: 3000000
}
func ExampleDeposit_inactive() {
	card := types.Card{Balance: 20_000_00, Active: false}
	Deposit(&card, 10_000_00)
	fmt.Println(card.Balance)

	// Output: 2000000
}
func ExampleDeposit_limit() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&card, 50_000_01)
	fmt.Println(card.Balance)

	// Output: 2000000
}
func ExampleAddBonus_normal() {
	card := types.Card{Balance: 20_000_00, MinBalance: 10_000_00, Active: true}
	AddBonus(&card, 3, 30, 100)
	fmt.Println(card.Balance)

	// Output: 2009000
}
func ExampleAddBonus_inactive() {
	card := types.Card{Balance: 20_000_00, MinBalance: 10_000_00, Active: false}
	AddBonus(&card, 3, 30, 100)
	fmt.Println(card.Balance)

	// Output: 2000000
}
func ExampleAddBonus_negative() {
	card := types.Card{Balance: -20_000_00, MinBalance: 10_000_00, Active: true}
	AddBonus(&card, 3, 30, 100)
	fmt.Println(card.Balance)

	// Output: -2000000
}
func ExampleAddBonus_limit() {
	card := types.Card{Balance: 20_000_00, MinBalance: 10_000_000_00, Active: true}
	AddBonus(&card, 3, 30, 100)
	fmt.Println(card.Balance)

	// Output: 2000000
}*/

func ExampleTotal()  {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active: true,
		},
		{
			Balance: 10_000_00,
			Active: false,
		},
		{
			Balance: 10_000_00,
			Active: true,
		},
	}

	fmt.Println(Total(cards))
	// Output: 2000000
}

func ExamplePaymentSources()  {
	cards := []types.Card{
		{
			PAN: "5058 2700 1013 8450",
			Balance: 10_000_00,
			Active: true,
		},
		{
			PAN: "5058 2702 8014 4303",
			Balance: 10_000_00,
			Active: false,
		},
		{
			PAN: "5058 2703 1246 1713",
			Balance: 10_000_00,
			Active: true,
		},
	}
	sources := PaymentSources(cards)
	for _, source := range sources {
		fmt.Println(source.Number)
	}
	// Output:
	// 5058 2700 1013 8450
	// 5058 2703 1246 1713
}