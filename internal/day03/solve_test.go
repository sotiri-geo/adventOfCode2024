package day03

import "testing"

func TestMultiply(t *testing.T) {
	t.Run("mul(4,3) returns 12", func(t *testing.T) {
		given := "mul(4,3)"
		got, _ := Multiply(given)
		want := 12

		assertMultiplyEquals(t, got, want)
	})

	t.Run("requires two numbers", func(t *testing.T) {
		given := "mul(,)"
		_, err := Multiply(given)

		assertError(t, err, ErrInsufficientMultiplyArgs)
	})
}

func assertMultiplyEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected error but did not get one")
	}

	// compare the underlining error type
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

}
