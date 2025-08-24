package day06

import (
	"errors"
	"testing"
)

func TestNewGuard(t *testing.T) {
	cases := []struct {
		Name     string
		InputMap [][]string
		Want     Guard
	}{
		{Name: "guard facing up", InputMap: [][]string{{".", "."}, {"#", "^"}}, Want: Guard{1, 1, Up, 1}},
		{Name: "guard facing right", InputMap: [][]string{{".", "."}, {">", "."}}, Want: Guard{1, 0, Right, 1}},
		{Name: "guard facing down", InputMap: [][]string{{".", "v"}, {".", "."}}, Want: Guard{0, 1, Down, 1}},
		{Name: "guard facing left", InputMap: [][]string{{".", "<"}, {".", "."}}, Want: Guard{0, 1, Left, 1}},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {

			got, _ := NewGuard(tt.InputMap)

			if *got != tt.Want {
				t.Errorf("got %v, want %v", *got, tt.Want)
			}
		})
	}

	t.Run("cannot find start position", func(t *testing.T) {
		inputMap := [][]string{{".", "."}}
		_, err := NewGuard(inputMap)

		assertError(t, err, ErrNotFoundGuardStartPosition)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("didn't get an error but required one")
	}

	if !errors.Is(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
