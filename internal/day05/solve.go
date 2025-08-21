package day05

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var ErrNoPageNumber = errors.New("no page number recorded as key.")

// Key is a page, values is an array of predecessor pages
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

func (p Predecessor) IsValidOrder(pageUpdates []int) (bool, error) {

	for idx, page := range pageUpdates {
		previousPages := pageUpdates[:idx]
		for _, previousPage := range previousPages {
			ok, _ := p.HasPredecessor(page, previousPage)
			if !ok {
				errorMessage := fmt.Sprintf("Failed with page: %d, previousPage: %d", page, previousPage)
				return false, errors.New(errorMessage)
			}
		}
	}
	return true, nil
}

func MiddleNumber(pageUpdates []int) int {
	middleIndex := len(pageUpdates) / 2
	return pageUpdates[middleIndex]
}

func NewPredecessor(orderingRules []string) Predecessor {
	pre := Predecessor{}
	for _, rule := range orderingRules {
		pre.Add(rule)
	}
	return pre
}

func Part1(predecessor Predecessor, pageUpdates [][]int) int {
	total := 0

	for _, pages := range pageUpdates {
		if ok, err := predecessor.IsValidOrder(pages); err == nil {
			if ok {
				total += MiddleNumber(pages)
			}
		}
	}
	return total
}
