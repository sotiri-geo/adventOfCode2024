package day06

import (
	"errors"
)

var (
	ErrNotFoundGuardStartPosition = errors.New("cannot find guards starting position")
)

type Direction int

// Iota is a counter starting from 0, 1...
const (
	Up Direction = iota
	Right
	Down
	Left
)

// State variables for Guard
// row, column, direction
type Guard struct {
	row           int
	column        int
	direction     Direction
	positionCount int
}

func NewGuard(inputMap [][]string) (*Guard, error) {
	directions := map[string]Direction{
		"^": Up,
		">": Right,
		"v": Down,
		"<": Left,
	}
	for ridx, row := range inputMap {
		for cidx, _ := range row {
			object := inputMap[ridx][cidx]
			direction, ok := directions[object]
			if ok {
				return &Guard{ridx, cidx, direction, 1}, nil
			}
		}
	}
	return nil, ErrNotFoundGuardStartPosition
}
