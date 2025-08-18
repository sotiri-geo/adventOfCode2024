package day03

import (
	"slices"
	"testing"
)

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

func TestExtractMultiply(t *testing.T) {
	extractionTests := map[string]struct {
		Input string
		Want  []string
	}{
		"extract one multiply expression":  {Input: "xmul(2,4)%&mul[3,7]!@^do_not", Want: []string{"mul(2,4)"}},
		"extract four multiply expression": {Input: "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", Want: []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"}},
	}

	for name, tt := range extractionTests {
		t.Run(name, func(t *testing.T) {
			got := ExtractMultiply(tt.Input)
			if !slices.Equal(got, tt.Want) {
				t.Errorf("got %v, want %v", got, tt.Want)
			}
		})
	}
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
