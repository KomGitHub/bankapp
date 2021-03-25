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
