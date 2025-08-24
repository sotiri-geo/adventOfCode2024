package day06

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewGuard(t *testing.T) {
	cases := []struct {
		Name     string
		InputMap [][]string
		Want     Guard
	}{
		{Name: "guard facing up", InputMap: [][]string{{".", "."}, {"#", "^"}}, Want: Guard{1, 1, Up, 1, true, [][]bool{{false, false}, {false, false}}}},
		{Name: "guard facing right", InputMap: [][]string{{".", "."}, {">", "."}}, Want: Guard{1, 0, Right, 1, true, [][]bool{{false, false}, {false, false}}}},
		{Name: "guard facing down", InputMap: [][]string{{".", "v"}, {".", "."}}, Want: Guard{0, 1, Down, 1, true, [][]bool{{false, false}, {false, false}}}},
		{Name: "guard facing left", InputMap: [][]string{{".", "<"}, {".", "."}}, Want: Guard{0, 1, Left, 1, true, [][]bool{{false, false}, {false, false}}}},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {

			got, err := NewGuard(tt.InputMap)

			if err != nil {
				t.Fatal("should not be an error")
			}

			assertGuardEqual(t, *got, tt.Want)
		})
	}
	t.Run("cannot find start position", func(t *testing.T) {
		inputMap := [][]string{{".", "."}}
		_, err := NewGuard(inputMap)

		assertError(t, err, ErrNotFoundGuardStartPosition)
	})
}

func TestGuardMoveForward(t *testing.T) {
	cases := []struct {
		name     string
		want     Guard
		inputMap [][]string
	}{
		{name: "moves forward 1 step", want: Guard{row: 0, column: 1, steps: 2, direction: Up, isPatrolling: true, visited: [][]bool{{false, false}, {false, false}}}, inputMap: [][]string{{".", "."}, {".", "^"}}},
		{name: "facing up against wall, rotate right 90 degrees", want: Guard{row: 1, column: 1, steps: 2, direction: Right, isPatrolling: true, visited: [][]bool{{false, false}, {false, false}}}, inputMap: [][]string{{"#", "."}, {"^", "."}}},
		{name: "facing left against wall, rotate right 90 degrees", want: Guard{row: 0, column: 1, steps: 2, direction: Up, isPatrolling: true, visited: [][]bool{{false, false}, {false, false}}}, inputMap: [][]string{{".", "."}, {"#", "<"}}},
		{name: "guard leaves the map", want: Guard{row: 0, column: 0, steps: 1, direction: Up, isPatrolling: false, visited: [][]bool{{false, false}, {false, false}}}, inputMap: [][]string{{"^", "."}, {".", "."}}},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGuard(tt.inputMap)
			t.Logf("Starts with Guard: %+v", *got)

			got.MoveFoward(tt.inputMap)

			if err != nil {
				t.Fatalf("should not have errored. Found: %v", err)
			}

			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("got %+v, want %+v", *got, tt.want)
			}

		})
	}
}

// func TestPart1(t *testing.T) {
//     t.Run("2x2 lab", func(t *testing.T) {
//         labInput := [][]string{{".", "."}, {".", "^"}}
//         got := Part1(labInput)
//         want := 2
//     })
// }

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("didn't get an error but required one")
	}

	if !errors.Is(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertGuardEqual(t testing.TB, got, want Guard) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
