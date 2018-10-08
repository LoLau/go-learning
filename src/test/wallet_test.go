package test

import (
	"testing"
	"fmt"
	"errors"
)

var insufficientFundError = errors.New("cannot withdraw, insufficient funds")

type Stringer interface {
	String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BIC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return insufficientFundError
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	assertError := func(t *testing.T, got, want error) {
		// Fatal会使程序退出
		if got == nil {
			t.Fatal("didnt get an error but wanted one.")
		}

		if got != want {
			t.Errorf("got %s, want %s.", got, want)
		}
	}

	assertNoError := func(t *testing.T, got error) {
		if got != nil {
			t.Errorf("got an error but didnt want one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startBalance := Bitcoin(20)
		wallet := Wallet{startBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startBalance)
		assertError(t, err, insufficientFundError)
	})
}





