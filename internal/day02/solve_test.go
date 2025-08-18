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

func TestIsDecreasing(t *testing.T) {
	t.Run("report is decreasing", func(t *testing.T) {
		given := []int{7, 6, 4, 2, 1}
		got := IsDecreasing(given)
		want := true

		if got != want {
			t.Errorf("got %v, want %v, given %v", got, want, given)
		}
	})

	t.Run("report is not decreasing", func(t *testing.T) {
		given := []int{7, 6, 8, 2, 1}
		got := IsDecreasing(given)
		want := false

		if got != want {
			t.Errorf("got %v, want %v, given %v", got, want, given)
		}
	})
}

func TestIsSafe(t *testing.T) {
	safeReportTest := map[string]struct {
		Report []int
		Want   bool
	}{
		"Increasing report is safe":                         {Report: []int{1, 3, 5, 6, 7}, Want: true},
		"Decreasing report is safe":                         {Report: []int{7, 6, 4, 2, 1}, Want: true},
		"Increase of 5 between 2 and 7 makes report unsafe": {Report: []int{1, 2, 7, 8, 9}, Want: false},
		"Decrease of 4 between 6 and 2 makes report unsafe": {Report: []int{9, 7, 6, 2, 1}, Want: false},
	}

	for name, tt := range safeReportTest {
		t.Run(name, func(t *testing.T) {
			got := IsSafe(tt.Report)

			if got != tt.Want {
				t.Errorf("got %v, want %v from report %v", got, tt.Want, tt.Report)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	t.Run("One report is safe", func(t *testing.T) {
		inputs := []string{"7 6 4 2 1"}
		got := Part1(inputs)
		want := 1

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Two out of six reports are safe", func(t *testing.T) {
		inputs := []string{"7 6 4 2 1", "1 2 7 8 9", "9 7 6 2 1", "1 3 2 4 5", "8 6 4 4 1", "1 3 6 7 9"}
		got := Part1(inputs)
		want := 2

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

// Part 2 tests

func TestIsSafeWithTolerance(t *testing.T) {
	safWithToleranceTest := map[string]struct {
		Report []int
		Want   bool
	}{
		"Report is safe when allowing for tolerance":    {Report: []int{1, 3, 2, 4, 5}, Want: true},
		"Report is unsafe after allowing for tolerance": {Report: []int{1, 2, 7, 8, 9}, Want: false},
	}

	for name, tt := range safWithToleranceTest {
		t.Run(name, func(t *testing.T) {
			got := IsSafeWithTolerance(tt.Report)

			if got != tt.Want {
				t.Errorf("got %v, want %v, given %+v", got, tt.Want, tt.Report)
			}
		})
	}
}
