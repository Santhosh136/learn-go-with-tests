package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat character given times", func(t *testing.T) {
		got := Repeat("s", 5)
		want := "sssss"

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}
