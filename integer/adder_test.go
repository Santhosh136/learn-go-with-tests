package integer

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("add two integers", func(t *testing.T) {
		sum := Add(24, 22)
		want := 46

		if sum != want {
			t.Errorf("want %d but got %d", want, sum)
		}
	})
}

func ExampleAdd() {
	sum := Add(2, 3)
	fmt.Println(sum)
	// Output: 5
}
