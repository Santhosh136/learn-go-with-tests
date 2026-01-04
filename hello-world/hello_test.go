package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("say hello with name", func(t *testing.T) {
		got := Hello("Santhosh", "")
		want := "Hello Santhosh"
		assert(t, got, want)
	})

	t.Run("say hello world when no name", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"
		assert(t, got, want)
	})

	t.Run("say hello in Tamil", func(t *testing.T) {
		got := Hello("Santhosh", "Tamil")
		want := "Vanakkam Santhosh"
		assert(t, got, want)
	})

	t.Run("say hello in Hindi", func(t *testing.T) {
		got := Hello("Santhosh", "Hindi")
		want := "Namaste Santhosh"
		assert(t, got, want)
	})
}

func assert(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
