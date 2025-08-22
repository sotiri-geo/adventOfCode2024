package day05

import (
	"reflect"
	"slices"
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

func TestGraph(t *testing.T) {
	t.Run("add edge to graph updates adjacency and indegree", func(t *testing.T) {
		pages := []int{1, 2, 3}
		graph := NewGraph(pages)

		graph.AddEdge(1, 2)

		wantAdj := map[int][]int{1: {2}}
		wantIndegree := map[int]int{1: 0, 2: 1, 3: 0}

		if !reflect.DeepEqual(graph.Adj, wantAdj) {
			t.Errorf("Adj: got %v, want %v", graph.Adj, wantAdj)
		}

		if !reflect.DeepEqual(graph.Indegree, wantIndegree) {
			t.Errorf("Indegree: got %v, want %v", graph.Indegree, wantIndegree)
		}
	})

	t.Run("extract 0 in degree pages", func(t *testing.T) {
		graph := Graph{
			Adj:      map[int][]int{1: {2}},
			Indegree: map[int]int{2: 1, 1: 0},
		}

		got := graph.ZeroIndegreePages()
		want := []int{1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("process a zero indegree page", func(t *testing.T) {
		// Find all new zero indegree pages to add
		graph := Graph{
			Adj:      map[int][]int{1: {2}},
			Indegree: map[int]int{2: 1, 1: 0},
		}

		got := graph.ProcessPage(1)
		want := []int{2}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("sort by page dependencies by top sort", func(t *testing.T) {
		graph := Graph{
			Adj:      map[int][]int{1: {2}},
			Indegree: map[int]int{2: 1, 1: 0},
		}

		got, _ := graph.TopSort()
		want := []int{1, 2}

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("found cycle in page dependencies", func(t *testing.T) {
		graph := Graph{
			Adj:      map[int][]int{1: {2}, 2: {1}},
			Indegree: map[int]int{2: 1, 1: 1},
		}

		_, err := graph.TopSort()

		assertError(t, err, ErrCycleFoundInPages)
	})
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
