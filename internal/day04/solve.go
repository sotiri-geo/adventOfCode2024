package day04

import (
	"strings"
)

const (
	XMAS       = "XMAS"
	xmasLength = len(XMAS)
)

type SearchXmas struct {
	matrix [][]string
}

func (s *SearchXmas) RightSearch(row, col int) bool {
	var buffer strings.Builder
	buffer.Grow(xmasLength)
	upperBound := min(col+xmasLength, len(s.matrix[0]))

	for i := col; i < upperBound; i++ {
		buffer.WriteString(s.matrix[row][i])
	}

	found := buffer.String()
	return found == XMAS
}

func (s *SearchXmas) LeftSearch(row, col int) bool {
	var buffer strings.Builder
	buffer.Grow(xmasLength)
	lowerBound := max(0, col-xmasLength)

	for i := col; i > lowerBound; i-- {
		buffer.WriteString(s.matrix[row][i])
	}

	found := buffer.String()
	return found == XMAS
}

func (s *SearchXmas) DownSearch(row, col int) bool {
	var buffer strings.Builder
	buffer.Grow(xmasLength)
	upperBound := min(row+xmasLength, len(s.matrix))

	for i := row; i < upperBound; i++ {
		buffer.WriteString(s.matrix[i][col])
	}

	found := buffer.String()
	return found == XMAS
}

func (s *SearchXmas) UpSearch(row, col int) bool {
	var buffer strings.Builder
	buffer.Grow(xmasLength)
	upperBound := max(row-xmasLength, 0)

	for i := row; i > upperBound; i-- {
		buffer.WriteString(s.matrix[i][col])
	}

	found := buffer.String()
	return found == XMAS
}

func (s *SearchXmas) UpRightSearch(row, col int) bool {
	var buffer strings.Builder
	buffer.Grow(xmasLength)

	for s.isValid(row, col, buffer) {
		buffer.WriteString(s.matrix[row][col])
		row--
		col++
	}

	found := buffer.String()
	return found == XMAS
}

func (s *SearchXmas) UpLeftSearch(row, col int) bool {
	var buffer strings.Builder

	for s.isValid(row, col, buffer) {
		buffer.WriteString(s.matrix[row][col])
		row--
		col--
	}

	found := buffer.String()
	return found == XMAS
}

func (s *SearchXmas) DownRightSearch(row, col int) bool {
	var buffer strings.Builder

	for s.isValid(row, col, buffer) {
		buffer.WriteString(s.matrix[row][col])
		row++
		col++
	}

	found := buffer.String()
	return found == XMAS
}

func (s *SearchXmas) DownLeftSearch(row, col int) bool {
	var buffer strings.Builder

	for s.isValid(row, col, buffer) {
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
