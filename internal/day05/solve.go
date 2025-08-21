package day05

import (
	"errors"
	"slices"
	"strconv"
	"strings"
)

var ErrNoPageNumber = errors.New("no page number recorded as key.")

type Predecessor map[int][]int

// Pointer reciever (p *Predecessor)
// Value reciever (p Predessor) <<< these types are references so value recievers fine
func (p Predecessor) Add(order string) {
	split := strings.Split(order, "|")
	from, _ := strconv.Atoi(split[0])
	to, _ := strconv.Atoi(split[1])
	p[to] = append(p[to], from)
}

func (p Predecessor) HasPredecessor(to, from int) (bool, error) {
	values, ok := p[to]

	if !ok {
		return false, ErrNoPageNumber
	}
	return slices.Contains(values, from), nil
}
