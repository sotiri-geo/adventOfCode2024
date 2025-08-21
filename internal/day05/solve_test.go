package day05

import (
	"reflect"
	"testing"
)

/*
<<< Ideas >>>

Create a map which shows for each key, which numbers must be prior to it

75,47,61,53,29
75|47, 75|61, 75|53, and 75|29.

We need to see for every page, the prior pages in the update are in the currents page
predecessors. If its not, then there exist one page thats out of order as part of update
*/

func TestPredecessor(t *testing.T) {
	t.Run("adds predessor value to collection", func(t *testing.T) {
		order := "75|47"
		got := Predecessor{}
		got.Add(order)
		want := Predecessor{
			47: []int{75},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Has predecessor value", func(t *testing.T) {
		pre := Predecessor{47: []int{75, 60}}
		to, from := 47, 75
		got, _ := pre.HasPredecessor(to, from)

		want := true

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Has no predecessor value", func(t *testing.T) {
		pre := Predecessor{47: []int{75, 60}}
		to, from := 47, 1
		got, _ := pre.HasPredecessor(to, from)

		want := false

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Has no key error", func(t *testing.T) {
		pre := Predecessor{47: []int{75, 60}}
		to, from := 10, 1
		_, got := pre.HasPredecessor(to, from)
		want := ErrNoPageNumber

		if got == nil {
			t.Fatal("did not raise no page number error")
		}

		if got != ErrNoPageNumber {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Has all prior pages as predecessors", func(t *testing.T) {
		pre := Predecessor{47: []int{70, 30}, 70: []int{30}}
		pageUpdates := []int{30, 70, 47}
		got, _ := pre.IsValidOrder(pageUpdates)
		want := true

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Works on example input", func(t *testing.T) {
		pre := Predecessor{47: []int{97, 75}, 61: []int{75, 47, 97}, 53: []int{47, 75, 61, 97}, 29: []int{75, 97, 53, 47, 61}}
		pageUpdates := []int{75, 47, 61, 53, 29}
		got, error := pre.IsValidOrder(pageUpdates)
		want := true

		t.Logf("got error: %v", error)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("init new map with defined predecessors", func(t *testing.T) {
		orderingRules := []string{"47|53", "97|13"}
		got := NewPredecessor(orderingRules)
		want := Predecessor{13: []int{97}, 53: []int{47}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestMiddleNumber(t *testing.T) {
	t.Run("odd length array", func(t *testing.T) {
		array := []int{3, 4, 5}
		got := MiddleNumber(array)
		want := 4

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("even length array", func(t *testing.T) {
		array := []int{7, 8, 9, 10}
		got := MiddleNumber(array)
		want := 9 // page number "just after" halfway point

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestParseUpdates(t *testing.T) {
	t.Run("parse an array of strings to 2D array of ints", func(t *testing.T) {
		input := []string{"1,2,3"}
		got := ParseUpdates(input)
		want := [][]int{{1, 2, 3}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestPart1(t *testing.T) {
	t.Run("Two page updates correct", func(t *testing.T) {
		pre := Predecessor{5: []int{1, 2, 3, 4}, 4: []int{1, 2, 3}, 3: []int{1, 2}, 2: []int{1}}
		pageUpdates := [][]int{{1, 2, 3}, {3, 2, 1}, {1, 2, 3, 4}}
		got := Part1(pre, pageUpdates)
		want := 5

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestAdjacencyList(t *testing.T) {
	cases := []struct {
		name          string
		pages         []int
		orderingRules []string
		want          AdjacencyList
	}{
		{name: "ordering rule pages subset of page updates", pages: []int{47, 53, 97, 13}, orderingRules: []string{"47|53", "97|13"}, want: AdjacencyList{47: []int{53}, 97: []int{13}}},
		{name: "ordering rule pages not a subset of page updates", pages: []int{47, 53, 97, 13}, orderingRules: []string{"47|53", "97|13", "67|18", "47|93"}, want: AdjacencyList{47: []int{53}, 97: []int{13}}},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAdjacencyList(tt.orderingRules, tt.pages)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInDegree(t *testing.T) {
	t.Run("init indegree map of pages", func(t *testing.T) {
		pages := []int{47, 53, 97, 13, 20}
		orderingRules := []string{"47|53", "97|13"}
		got := NewInDegree(orderingRules, pages)
		want := InDegree{53: 1, 13: 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("find zero indegree page", func(t *testing.T) {
		pages := []int{53, 47, 13}
		// Zero degree nodes are not stored, search via pages to see missing node from InDegree
		inDegree := InDegree{53: 1, 13: 1}
		got, _ := inDegree.FirstZeroInDegree(pages)
		want := 47

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("multiple zero in degree pages", func(t *testing.T) {
		pages := []int{53, 47, 13, 11}
		// Zero degree nodes are not stored, search via pages to see missing node from InDegree
		inDegree := InDegree{53: 1, 13: 1}
		_, err := inDegree.FirstZeroInDegree(pages)

		assertError(t, err, ErrMultipleZeroInDegreePages)
	})

	t.Run("no zero in degree pages", func(t *testing.T) {
		pages := []int{53, 13}
		// Zero degree nodes are not stored, search via pages to see missing node from InDegree
		inDegree := InDegree{53: 1, 13: 1}
		_, err := inDegree.FirstZeroInDegree(pages)

		assertError(t, err, ErrNoZeroInDegreePages)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("did not get an error but wanted one")
	}
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
