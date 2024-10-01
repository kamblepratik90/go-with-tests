package main

import (
	"errors"
	"fmt"
	"testing"
)

// create new types from existing ones
type Bitcoin int

type Stringer interface {
	String() string
}

// This interface is defined in the fmt package and lets you define how your type
// is printed when used with the %s format string in prints.

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	// balance int
	balance Bitcoin
}

// func (w *Wallet) Deposit(a int) {
func (w *Wallet) Deposit(a Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p bal: %s \n", &w.balance, w.Balance())
	w.balance += a
	fmt.Printf("balance after Deposit is bal: %s \n", w.Balance())
}

// func (w *Wallet) Balance() int {
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(a Bitcoin) error {
	fmt.Printf("address of balance in Withdraw is %p  bal: %s \n", &w.balance, w.Balance())

	if w.Balance() < a {
		return ErrInsufficientFunds
	}

	w.balance -= a
	fmt.Printf("balance after Withdraw is bal: %s \n", w.Balance())

	return nil
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s and want %s \n", got, want)
		}
	}

	assertNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didn't want one")
		}
	}

	assertError := func(t testing.TB, got, want error) {
		t.Helper()

		if got == nil {
			t.Error("Wanted an err but dint got any")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		// wallet.Deposit(10)
		wallet.Deposit(Bitcoin(10))

		fmt.Printf("address of balance in test is %p \n", &wallet.balance)
		// The %p placeholder prints memory addresses in base 16 notation with leading 0xs and the  escape character prints a new line.

		// got := wallet.Balance()
		// // expect := 10
		// expect := Bitcoin(10)

		// if got != expect {
		// 	// t.Errorf("got %d and expect %d", got, expect)
		// 	t.Errorf("got %s want %s", got, expect)
		// }

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		// wallet := Wallet{}
		wallet := Wallet{balance: Bitcoin(20)}

		// // deposit
		// wallet.Deposit(Bitcoin(20))

		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		fmt.Printf("address of balance in test is %p \n", &wallet.balance)
		// The %p placeholder prints memory addresses in base 16 notation with leading 0xs and the  escape character prints a new line.

		// got := wallet.Balance()
		// expect := Bitcoin(10)

		// if got != expect {
		// 	// t.Errorf("got %d and expect %d", got, expect)
		// 	t.Errorf("got %s want %s", got, expect)
		// }

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		initBal := Bitcoin(20)
		wallet := Wallet{balance: initBal}

		err := wallet.Withdraw(Bitcoin(200))

		assertError(t, err, ErrInsufficientFunds)

		assertBalance(t, wallet, initBal)

		// if err == nil {
		// 	t.Errorf("Wanted an err but dint got any")
		// }

	})

}
