package day04

import (
	"testing"
)

func TestRightSearchXmas(t *testing.T) {
	searchXmas := SearchXmas{matrix: [][]string{{"X", "M", "A", "S", "X"}}}
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
	leftSearchTests := []struct {
		name   string
		row    int
		col    int
		matrix [][]string
		want   bool
	}{
		{name: "Found XMAS", row: 0, col: 3, matrix: [][]string{{"S", "A", "M", "X"}}, want: true},
		{name: "Cannot find XMAS", row: 0, col: 2, matrix: [][]string{{"S", "A", "M", "X", "X"}}, want: false},
		{name: "Out of bounds", row: 0, col: 2, matrix: [][]string{{"S", "S", "A", "M", "X"}}, want: false},
	}

	for _, tt := range leftSearchTests {
		t.Run(tt.name, func(t *testing.T) {
			searchXmas := SearchXmas{matrix: tt.matrix}
			got := searchXmas.LeftSearch(tt.row, tt.col)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDownSearchXmas(t *testing.T) {
	searchXmas := SearchXmas{matrix: [][]string{{"X", "S", "A"}, {"M", "S", "A"}, {"A", "S", "A"}, {"S", "S", "A"}, {"X", "S", "A"}}}
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
	upSearchTests := []struct {
		name   string
		row    int
		col    int
		matrix [][]string
		want   bool
	}{
		{name: "Found XMAS", row: 3, col: 0, matrix: [][]string{{"S"}, {"A"}, {"M"}, {"X"}}, want: true},
		{name: "Cannot find XMAS", row: 3, col: 0, matrix: [][]string{{"A"}, {"S"}, {"A"}, {"M"}, {"X"}}, want: false},
		{name: "Out of bounds", row: 2, col: 0, matrix: [][]string{{"S"}, {"A"}, {"M"}, {"X"}}, want: false},
	}

	for _, tt := range upSearchTests {
		t.Run(tt.name, func(t *testing.T) {
			searchXmas := SearchXmas{matrix: tt.matrix}
			got := searchXmas.UpSearch(tt.row, tt.col)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpRightSearchXmas(t *testing.T) {
	searchXmas := SearchXmas{matrix: [][]string{{"X", "S", "A", "X"}, {"S", "S", "A", "S"}, {"A", "S", "A", "S"}, {"M", "M", "A", "S"}, {"X", "S", "A", "M"}}}
	upRightSearchTests := []struct {
		name string
		row  int
		col  int
		want bool
	}{
		{name: "Found XMAS", row: 4, col: 0, want: true},
		{name: "Cannot find XMAS", row: 3, col: 0, want: false},
		{name: "Out of bounds", row: 2, col: 0, want: false},
	}

	for _, tt := range upRightSearchTests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchXmas.UpRightSearch(tt.row, tt.col)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpLeftSearchXmas(t *testing.T) {
	searchXmas := SearchXmas{matrix: [][]string{{"X", "S", "A", "X"}, {"S", "S", "A", "S"}, {"A", "A", "A", "S"}, {"M", "M", "M", "S"}, {"X", "S", "A", "X"}}}
	upLeftSearchTests := []struct {
		name string
		row  int
		col  int
		want bool
	}{
		{name: "Found XMAS", row: 4, col: 3, want: true},
		{name: "Cannot find XMAS", row: 3, col: 3, want: false},
		{name: "Out of bounds", row: 2, col: 0, want: false},
	}

	for _, tt := range upLeftSearchTests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchXmas.UpLeftSearch(tt.row, tt.col)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDownRightSearchXmas(t *testing.T) {
	searchXmas := SearchXmas{matrix: [][]string{{"X", "S", "A", "X"}, {"S", "M", "A", "S"}, {"A", "A", "A", "S"}, {"M", "M", "M", "S"}, {"X", "S", "A", "X"}}}
	downRightSearchTests := []struct {
		name string
		row  int
		col  int
		want bool
	}{
		{name: "Found XMAS", row: 0, col: 0, want: true},
		{name: "Cannot find XMAS", row: 1, col: 0, want: false},
		{name: "Out of bounds", row: 3, col: 0, want: false},
	}

	for _, tt := range downRightSearchTests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchXmas.DownRightSearch(tt.row, tt.col)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDownLeftSearchXmas(t *testing.T) {
	searchXmas := SearchXmas{matrix: [][]string{{"X", "S", "A", "X"}, {"S", "M", "M", "S"}, {"A", "A", "A", "S"}, {"S", "M", "M", "S"}, {"X", "S", "A", "X"}}}
	downLeftSearchTests := []struct {
		name string
		row  int
		col  int
		want bool
	}{
		{name: "Found XMAS", row: 0, col: 3, want: true},
		{name: "Cannot find XMAS", row: 1, col: 3, want: false},
		{name: "Out of bounds", row: 3, col: 0, want: false},
	}

	for _, tt := range downLeftSearchTests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchXmas.DownLeftSearch(tt.row, tt.col)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindXmas(t *testing.T) {
	t.Run("Find 2 XMAS directionally right and down from current position", func(t *testing.T) {
		searchXmas := SearchXmas{matrix: [][]string{{"X", "M", "A", "S"}, {"M", "X", "X", "X"}, {"A", "X", "X", "X"}, {"S", "X", "X", "X"}, {"X", "X", "X", "X"}}}
		got := searchXmas.Find(0, 0)
		want := 2

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestPart1(t *testing.T) {
	t.Run("Has two XMAS in search puzzle", func(t *testing.T) {
		puzzle := [][]string{{"X", "M", "A", "S"}, {"M", "A", "S", "X"}, {"A", "S", "X", "S"}, {"S", "X", "M", "A"}}
		got := Part1(puzzle)
		want := 2

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Has XMAS occur 18 times. Example input.", func(t *testing.T) {
		input := []string{
			"MMMSXXMASM",
			"MSAMXMSMSA",
			"AMXSXMAAMM",
			"MSAMASMSMX",
			"XMASAMXAMM",
			"XXAMMXXAMA",
			"SMSMSASXSS",
			"SAXAMASAAA",
			"MAMMMXMMMM",
			"MXMXAXMASX",
		}
		puzzle := To2DMatrix(input)
		got := Part1(puzzle)
		want := 18

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestForwardMas(t *testing.T) {
	t.Run("found MAS", func(t *testing.T) {
		searchMas := SearchMas{matrix: [][]string{{"X", "M", "S"}, {"X", "A", "S"}, {"M", "M", "S"}}}
		got := searchMas.HasForward(1, 1)
		want := true

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("cannot find MAS", func(t *testing.T) {
		searchMas := SearchMas{matrix: [][]string{{"X", "M", "X"}, {"X", "A", "S"}, {"M", "M", "S"}}}
		got := searchMas.HasForward(1, 1)
		want := false

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("out of bounds", func(t *testing.T) {
		searchMas := SearchMas{matrix: [][]string{{"X", "M", "X"}, {"X", "A", "S"}, {"M", "M", "S"}}}
		got := searchMas.HasForward(0, 0)
		want := false

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestBackwardMas(t *testing.T) {
	t.Run("found MAS", func(t *testing.T) {
		searchMas := SearchMas{matrix: [][]string{{"M", "M", "X"}, {"X", "A", "S"}, {"M", "M", "S"}}}
		got := searchMas.HasBackward(1, 1)
		want := true

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("cannot find MAS", func(t *testing.T) {
		searchMas := SearchMas{matrix: [][]string{{"X", "M", "X"}, {"X", "A", "S"}, {"M", "M", "S"}}}
		got := searchMas.HasBackward(1, 1)
		want := false

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("out of bounds", func(t *testing.T) {
		searchMas := SearchMas{matrix: [][]string{{"X", "M", "X"}, {"X", "A", "S"}, {"M", "M", "S"}}}
		got := searchMas.HasBackward(0, 0)
		want := false

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
