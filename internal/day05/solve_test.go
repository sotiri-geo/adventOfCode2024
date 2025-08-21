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
}
