package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit money to the wallet", func(t *testing.T) {

		wallet := Wallet{}
		wallet.Deposit(Bitcoin(100))

		balance := wallet.Balance()
		want := Bitcoin(100)

		if balance != want {
			t.Errorf("got %s but want %s", balance, want)
		}

	})

}
