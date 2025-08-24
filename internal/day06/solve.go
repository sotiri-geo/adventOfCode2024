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

const Wall = "#"

type Position struct {
	row    int
	column int
}

// State variables for Guard
// row, column, direction
type Guard struct {
	row       int
	column    int
	direction Direction
	steps     int // starts at 1
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

func (g *Guard) MoveFoward(inputMap [][]string) {
	nextPosition := g.getNextPosition()

	if inputMap[nextPosition.row][nextPosition.column] == Wall {
		// We need to rotate 90 degrees then recursively call Move forward
		g.turn()
		g.MoveFoward(inputMap)
	} else {
		g.row = nextPosition.row
		g.column = nextPosition.column
		g.steps++
	}
}

func (g *Guard) getNextPosition() Position {
	switch g.direction {
	case Up:
		return Position{g.row - 1, g.column}
	case Right:
		return Position{g.row, g.column + 1}
	case Down:
		return Position{g.row + 1, g.column}
	case Left:
		return Position{g.row, g.column - 1}
	}
	return Position{g.row, g.column}
}

func (g *Guard) turn() {
	g.direction = (g.direction + 1) % 4
}
