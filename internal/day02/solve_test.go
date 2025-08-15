package day02

import "testing"

// We need to work out if a report is safe
// This is decided if a report which is an array of ints is either increasing or decreasing at a rate
// between 1 and 3 inclusive

func TestIncreaseWithinBounds(t *testing.T) {
	increasingTests := map[string]struct {
		Start int
		End   int
		Want  bool
	}{
		"Within bounds of 1 to 3":  {Start: 1, End: 3, Want: true},
		"Outside bounds of 1 to 3": {Start: 0, End: 5, Want: false},
		"Strictly decreasing":      {Start: 5, End: 3, Want: false},
		"Constant":                 {Start: 2, End: 2, Want: false},
	}

	for name, tt := range increasingTests {
		t.Run(name, func(t *testing.T) {
			got := IncreaseWithinBounds(tt.Start, tt.End)

			if got != tt.Want {
				t.Errorf("got %v, want %v, given start %d and end %d ", got, tt.Want, tt.Start, tt.End)
			}
		})
	}
}

func TestIsIncreasing(t *testing.T) {
	t.Run("report is increasing", func(t *testing.T) {
		given := []int{1, 3, 6, 7, 9}
		got := IsIncreasing(given)
		want := true

		if got != want {
			t.Errorf("got %v, want %v, given %+v", got, want, given)
		}
	})

	t.Run("report is not increasing", func(t *testing.T) {
		given := []int{1, 3, 6, 4, 9}
		got := IsIncreasing(given)
		want := false

		if got != want {
			t.Errorf("got %v, want %v, given %v", got, want, given)
		}
	})
}

func TestDecreaseWithinBounds(t *testing.T) {
	decreasingTests := map[string]struct {
		Start int
		End   int
		Want  bool
	}{
		"Within bounds of 1 to 3":  {Start: 3, End: 1, Want: true},
		"Outside bounds of 1 to 3": {Start: 5, End: 0, Want: false},
		"Strictly increasing":      {Start: 3, End: 5, Want: false},
		"Constant":                 {Start: 2, End: 2, Want: false},
	}

	for name, tt := range decreasingTests {
		t.Run(name, func(t *testing.T) {
			got := DecreaseWithinBounds(tt.Start, tt.End)

			if got != tt.Want {
				t.Errorf("got %v, want %v, given start %d and end %d ", got, tt.Want, tt.Start, tt.End)
			}
		})
	}
}
