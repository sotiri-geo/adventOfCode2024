package day01

import "testing"

func TestDistance(t *testing.T) {
	distanceTests := map[string]struct {
		From int
		To   int
		Want int
	}{
		"From 4 to 6 should be 2": {From: 2, To: 6, Want: 4},
		"From 5 to 2 should be 3": {From: 5, To: 2, Want: 3},
	}

	for name, tt := range distanceTests {
		t.Run(name, func(t *testing.T) {
			got := Distance(tt.From, tt.To)

			if got != tt.Want {
				t.Errorf("got %d, want %d", got, tt.Want)
			}
		})
	}
}

// Another vertical slice, we need to sum across an array of ints

func TestSum(t *testing.T) {

	t.Run("Sum a non empty array of distances", func(t *testing.T) {
		distances := []int{2, 3, 4}
		got := Sum(distances)
		want := 9

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestSplit(t *testing.T) {
	parsedString, _ := split("4   9")
	got := Row{Left: parsedString.Left, Right: parsedString.Right}
	want := Row{Left: 4, Right: 9}

	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

// E2E test or Acceptance tests
func TestPart1(t *testing.T) {

	t.Run("A single pair of ids to compare", func(t *testing.T) {
		inputs := []string{"4   9"}
		got := Part1(inputs)
		want := 5

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Multiple pair of ids to compare", func(t *testing.T) {
		inputs := []string{"1   3", "4   9"}
		got := Part1(inputs)
		want := 7

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
