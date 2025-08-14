package day01

import (
	"maps"
	"testing"
)

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

// Part 2 tests

func TestCounter(t *testing.T) {

	t.Run("Creates a counter from a slice of ints", func(t *testing.T) {
		ids := []int{1, 1, 4, 5}
		got := Counter(ids)
		want := map[int]int{
			1: 2,
			4: 1,
			5: 1,
		}

		if !maps.Equal(got, want) {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})
}

func TestSimilarityScore(t *testing.T) {

	t.Run("computes similarity score with 3 ids", func(t *testing.T) {
		leftIds := []int{1, 3, 3}
		rightIds := []int{2, 3, 9}

		got := SimilarityScore(leftIds, rightIds)
		want := 6

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("computes basic similarity score with 5 ids", func(t *testing.T) {
		leftIds := []int{1, 2, 3, 7, 8}
		rightIds := []int{2, 2, 4, 5, 3}

		got := SimilarityScore(leftIds, rightIds)
		want := 7

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestPart2(t *testing.T) {

	t.Run("generates part 2 score for example case", func(t *testing.T) {
		inputs := []string{"3   4", "4   3", "2   5", "1   3", "3   9", "3   3"}
		got := Part2(inputs)
		want := 31

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}

	})
}
