package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

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

	t.Run("Withdraw money with insufficient balance", func(t *testing.T) {
		initialBalance := Bitcoin(100)
		wallet := Wallet{balance: initialBalance}
		err := wallet.Withdraw(Bitcoin(200))

		assertError(t, err, ErrorInsufficientBalance)
		assertBalance(t, wallet, initialBalance)
	})

}

func assertBalance(t testing.TB, w Wallet, want Bitcoin) {
	t.Helper()
	balance := w.Balance()
	if balance != want {
		t.Errorf("got %s but want %s", balance, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("didn't get error but wanted one")
	}

	if got.Error() != want.Error() {
		t.Errorf("got %q, want %q", got, want)
	}
}
