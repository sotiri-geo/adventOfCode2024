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
	row       int
	column    int
	direction Direction
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
		// g.MoveFoward(inputMap)
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
		return Position{g.row - 1, g.column, g.direction}
	case Right:
		return Position{g.row, g.column + 1, g.direction}
	case Down:
		return Position{g.row + 1, g.column, g.direction}
	case Left:
		return Position{g.row, g.column - 1, g.direction}
	}
	return Position{g.row, g.column, g.direction}
}

func (g *Guard) turnClockwise() {
	g.direction = (g.direction + 1) % 4
}

func (g *Guard) withinBoundary(position Position, inputMap [][]string) bool {
	rowLength := len(inputMap)
	columnLength := len(inputMap[0])

	return position.row >= 0 && position.row < rowLength && position.column >= 0 && position.column < columnLength
}

func Part1(inputMap [][]string) int {
	guard, _ := NewGuard(inputMap)
	seen := make(map[Position]struct{})

	initPosition := Position{row: guard.row, column: guard.column, direction: guard.direction}
	seen[initPosition] = struct{}{}

	for guard.isPatrolling {
		guard.MoveFoward(inputMap)
		newPosition := Position{row: guard.row, column: guard.column, direction: guard.direction}
		if _, ok := seen[newPosition]; ok {
			// prevents loops
			return guard.steps
		}
	}

	return guard.steps
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

func HasLoop(inputMap [][]string) bool {
	// We need to get back to a position we've seen
	seen := make(map[Position]struct{})
	guard, _ := NewGuard(inputMap)

	for guard.isPatrolling {
		currentPosition := Position{row: guard.row, column: guard.column, direction: guard.direction}
		if _, exists := seen[currentPosition]; exists {
			return true // loop detected
		}

		seen[currentPosition] = struct{}{}
		guard.MoveFoward(inputMap)
	}

	return false
}

func Part2(inputMap [][]string) int {
	count := 0

	// Get guards starting position to avoid placing obstruction there
	guard, _ := NewGuard(inputMap)
	startRow, startCol := guard.row, guard.column

	for i := range inputMap {
		for j := range inputMap[i] {
			if inputMap[i][j] == "." && !(i == startRow && j == startCol) {
				// attempt an obstruction
				inputMap[i][j] = "#"
				if HasLoop(inputMap) {
					count++
				}
				// Restore original value
				inputMap[i][j] = "."
			}
		}
	}
	return count
}
