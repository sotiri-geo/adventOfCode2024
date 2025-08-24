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
	row          int
	column       int
	direction    Direction
	steps        int  // starts at 1
	isPatrolling bool // guard in map
	visited      [][]bool
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
				visited := newBoolGridFrom(inputMap)
				visited[ridx][cidx] = true
				return &Guard{ridx, cidx, direction, 1, true, visited}, nil
			}
		}
	}
	return nil, ErrNotFoundGuardStartPosition
}

func newBoolGridFrom[T any](src [][]T) [][]bool {
	dst := make([][]bool, len(src))
	for i := range src {
		dst[i] = make([]bool, len(src[i]))
	}
	return dst
}

func (g *Guard) MoveFoward(inputMap [][]string) {
	nextPosition := g.getNextPosition()

	if !g.withinBoundary(nextPosition, inputMap) {
		g.isPatrolling = false
	} else if inputMap[nextPosition.row][nextPosition.column] == Wall {
		// We cannot walk through a wall
		g.turnClockwise()
		g.MoveFoward(inputMap)
	} else {
		g.row = nextPosition.row
		g.column = nextPosition.column
		// Only increment step if not visited before
		if !g.visited[g.row][g.column] {
			g.steps++
		}

		g.visited[g.row][g.column] = true
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

func (g *Guard) turnClockwise() {
	g.direction = (g.direction + 1) % 4
}

func (g *Guard) withinBoundary(position Position, inputMap [][]string) bool {
	rowLength := len(inputMap)
	columnLength := len(inputMap[0])

	return position.row >= 0 && position.row < rowLength && position.column >= 0 && position.column < columnLength
}

func Part1(labInput [][]string) int {
	guard, _ := NewGuard(labInput)

	for guard.isPatrolling {
		guard.MoveFoward(labInput)
	}

	return guard.steps
}
