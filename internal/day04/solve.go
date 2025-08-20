package day04

import (
	"strings"
)

const (
	XMAS       = "XMAS"
	xmasLength = len(XMAS)
)

// All we need to implement an interface in Go is just implement all methods on interface
type Search interface {
	Find() int
}

type SearchXmas struct {
	matrix [][]string
}

func (s *SearchXmas) RightSearch(row, col int) bool {
	var buffer strings.Builder
	buffer.Grow(xmasLength)

	for i := col; i < len(s.matrix[0]); i++ {
		if buffer.Len() == xmasLength {
			break
		}
		buffer.WriteString(s.matrix[row][i])
	}

	found := buffer.String()
	return found == XMAS
}

func (s *SearchXmas) LeftSearch(row, col int) bool {
	var buffer strings.Builder
	buffer.Grow(xmasLength)

	for i := col; i >= 0; i-- {
		if buffer.Len() == xmasLength {
			break
		}
		buffer.WriteString(s.matrix[row][i])
	}

	found := buffer.String()
	return found == XMAS
}

func (s *SearchXmas) DownSearch(row, col int) bool {
	var buffer strings.Builder
	buffer.Grow(xmasLength)

	for i := row; i < len(s.matrix); i++ {
		if buffer.Len() == xmasLength {
			break
		}
		buffer.WriteString(s.matrix[i][col])
	}

	found := buffer.String()
	return found == XMAS
}

func (s *SearchXmas) UpSearch(row, col int) bool {
	var buffer strings.Builder
	buffer.Grow(xmasLength)

	for i := row; i >= 0; i-- {
		if buffer.Len() == xmasLength {
			break
		}
		buffer.WriteString(s.matrix[i][col])
	}

	found := buffer.String()
	return found == XMAS
}

func (s *SearchXmas) UpRightSearch(row, col int) bool {
	var buffer strings.Builder
	buffer.Grow(xmasLength)

	for s.isValid(row, col, buffer) {
		if buffer.Len() == xmasLength {
			break
		}
		buffer.WriteString(s.matrix[row][col])
		row--
		col++
	}

	found := buffer.String()
	return found == XMAS
}

func (s *SearchXmas) UpLeftSearch(row, col int) bool {
	var buffer strings.Builder
	buffer.Grow(xmasLength)

	for s.isValid(row, col, buffer) {
		if buffer.Len() == xmasLength {
			break
		}
		buffer.WriteString(s.matrix[row][col])
		row--
		col--
	}

	found := buffer.String()
	return found == XMAS
}

func (s *SearchXmas) DownRightSearch(row, col int) bool {
	var buffer strings.Builder
	buffer.Grow(xmasLength)

	for s.isValid(row, col, buffer) {
		if buffer.Len() == xmasLength {
			break
		}
		buffer.WriteString(s.matrix[row][col])
		row++
		col++
	}

	found := buffer.String()
	return found == XMAS
}

func (s *SearchXmas) DownLeftSearch(row, col int) bool {
	var buffer strings.Builder
	buffer.Grow(xmasLength)

	for s.isValid(row, col, buffer) {
		if buffer.Len() == xmasLength {
			break
		}
		buffer.WriteString(s.matrix[row][col])
		row++
		col--
	}

	found := buffer.String()
	return found == XMAS
}

func (s *SearchXmas) isValid(row, col int, buffer strings.Builder) bool {
	rowInRange := row >= 0 && row < len(s.matrix)
	colInRange := col >= 0 && col < len(s.matrix[0])
	bufferCap := buffer.Len() < xmasLength

	return rowInRange && colInRange && bufferCap
}

func (s *SearchXmas) Find(row, col int) int {
	count := 0

	// 4 directional
	if s.UpSearch(row, col) {
		count++
	}
	if s.RightSearch(row, col) {
		count++
	}
	if s.DownSearch(row, col) {
		count++
	}
	if s.LeftSearch(row, col) {
		count++
	}
	// 4 diagonal
	if s.UpRightSearch(row, col) {
		count++
	}
	if s.UpLeftSearch(row, col) {
		count++
	}
	if s.DownLeftSearch(row, col) {
		count++
	}
	if s.DownRightSearch(row, col) {
		count++
	}
	return count
}

func Part1(puzzle [][]string) int {
	searchXmas := SearchXmas{matrix: puzzle}
	count := 0

	for row := 0; row < len(puzzle); row++ {
		for col := 0; col < len(puzzle[0]); col++ {
			count += searchXmas.Find(row, col)
		}
	}
	return count
}

func To2DMatrix(input []string) [][]string {
	output := make([][]string, len(input))

	for row := 0; row < len(input); row++ {
		// allocate each row with a slice of size len(input[row])
		output[row] = make([]string, len(input[row]))
		for col := 0; col < len(input[row]); col++ {
			output[row][col] = string(input[row][col])
		}
	}
	return output
}

// Part 2

type SearchMas struct {
	matrix [][]string
}

func (s *SearchMas) HasForward(row, col int) bool {
	current := s.matrix[row][col]

	if current != "A" {
		return false
	}

	straight := s.hasCharacterM(row+1, col-1) && s.hasCharacterS(row-1, col+1)
	reverse := s.hasCharacterS(row+1, col-1) && s.hasCharacterM(row-1, col+1)

	return straight || reverse
}

func (s *SearchMas) HasBackward(row, col int) bool {
	current := s.matrix[row][col]

	if current != "A" {
		return false
	}

	straight := s.hasCharacterM(row-1, col-1) && s.hasCharacterS(row+1, col+1)
	reverse := s.hasCharacterS(row-1, col-1) && s.hasCharacterM(row+1, col+1)

	return straight || reverse
}

func (s *SearchMas) isValid(row, col int) bool {
	rowInRange := row >= 0 && row < len(s.matrix)
	colInRange := col >= 0 && col < len(s.matrix[0])

	return rowInRange && colInRange
}

func (s *SearchMas) hasCharacterS(row, col int) bool {
	return s.isValid(row, col) && s.matrix[row][col] == "S"
}

func (s *SearchMas) hasCharacterM(row, col int) bool {
	return s.isValid(row, col) && s.matrix[row][col] == "M"
}

func (s *SearchMas) Find(row, col int) int {
	if s.HasForward(row, col) && s.HasBackward(row, col) {
		return 1
	}
	return 0
}

func Part2(puzzle [][]string) int {
	searchMas := SearchMas{matrix: puzzle}
	count := 0

	for row := 0; row < len(puzzle); row++ {
		for col := 0; col < len(puzzle[0]); col++ {
			result := searchMas.Find(row, col)
			count += result
		}
	}
	return count
}
