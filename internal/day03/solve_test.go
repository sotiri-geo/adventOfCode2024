package day03

import "testing"

func TestMultiply(t *testing.T) {
	t.Run("mul(4,3) returns 12", func(t *testing.T) {
		given := "mul(4,3)"
		got, _ := Multiply(given)
		want := 12

		if got != want {
			t.Errorf("got %d, want %d, given %s", got, want, given)
		}
	})

	t.Run("requires two numbers", func(t *testing.T) {
		given := "mul(,)"
		_, err := Multiply(given)

		if err == nil {
			t.Fatal("Expected error but did not get one")
		}

		if err != ErrInsufficientMultiplyArgs {
			t.Errorf("got %s, want %s", err, ErrInsufficientMultiplyArgs)
		}

	})
}
