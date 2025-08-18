package day04

import (
	"testing"
)

func TestRightSearchXmas(t *testing.T) {
	searchXmas := SearchXmas{matrix: [][]string{{"X", "M", "A", "S", "X"}}, Count: 0}
	rightSearchTests := []struct {
		name   string
		row    int
		col    int
		matrix [][]string
		want   bool
	}{
		{name: "Found XMAS", row: 0, col: 0, want: true},
		{name: "Cannot find XMAS", row: 0, col: 1, want: false},
		{name: "Out of bounds", row: 0, col: 3, want: false},
	}

	for _, tt := range rightSearchTests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchXmas.RightSearch(tt.row, tt.col)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLeftSearchXmas(t *testing.T) {
	searchXmas := SearchXmas{matrix: [][]string{{"X", "S", "A", "M", "X"}}, Count: 0}
	leftSearchTests := []struct {
		name string
		row  int
		col  int
		want bool
	}{
		{name: "Found XMAS", row: 0, col: 4, want: true},
		{name: "Cannot find XMAS", row: 0, col: 3, want: false},
		{name: "Out of bounds", row: 0, col: 2, want: false},
	}

	for _, tt := range leftSearchTests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchXmas.LeftSearch(tt.row, tt.col)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDownSearchXmas(t *testing.T) {
	searchXmas := SearchXmas{matrix: [][]string{{"X", "S", "A"}, {"M", "S", "A"}, {"A", "S", "A"}, {"S", "S", "A"}, {"X", "S", "A"}}, Count: 0}
	downSearchTests := []struct {
		name string
		row  int
		col  int
		want bool
	}{
		{name: "Found XMAS", row: 0, col: 0, want: true},
		{name: "Cannot find XMAS", row: 1, col: 0, want: false},
		{name: "Out of bounds", row: 4, col: 0, want: false},
	}

	for _, tt := range downSearchTests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchXmas.DownSearch(tt.row, tt.col)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpSearchXmas(t *testing.T) {
	searchXmas := SearchXmas{matrix: [][]string{{"X", "S", "A"}, {"S", "S", "A"}, {"A", "S", "A"}, {"M", "S", "A"}, {"X", "S", "A"}}, Count: 0}
	upSearchTests := []struct {
		name string
		row  int
		col  int
		want bool
	}{
		{name: "Found XMAS", row: 4, col: 0, want: true},
		{name: "Cannot find XMAS", row: 3, col: 0, want: false},
		{name: "Out of bounds", row: 2, col: 0, want: false},
	}

	for _, tt := range upSearchTests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchXmas.UpSearch(tt.row, tt.col)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
