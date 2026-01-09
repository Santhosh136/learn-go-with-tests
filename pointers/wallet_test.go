package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, w Wallet, want Bitcoin) {
		t.Helper()
		balance := w.Balance()
		if balance != want {
			t.Errorf("got %s but want %s", balance, want)
		}
	}

	t.Run("Deposit money to the wallet", func(t *testing.T) {

		wallet := Wallet{}
		wallet.Deposit(Bitcoin(100))

		want := Bitcoin(100)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw money from the wallet", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(200)}
		wallet.Withdraw(Bitcoin(100))

		want := Bitcoin(100)
		assertBalance(t, wallet, want)
	})

}
