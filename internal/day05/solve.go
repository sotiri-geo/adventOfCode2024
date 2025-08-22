package day05

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/sotiri-geo/adventOfCode2024/internal/day01"
)

var (
	ErrNoPageNumber              = errors.New("no page number recorded as key.")
	ErrMultipleZeroInDegreePages = errors.New("found multiple zero indegree pages")
	ErrNoZeroInDegreePages       = errors.New("Cannot find zero indegree page")
)

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

func ParseUpdates(updates []string) [][]int {
	parsed := make([][]int, len(updates))

	for idx, row := range updates {
		parsed[idx] = day01.ToInts(strings.Split(row, ","))
	}
	return parsed
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

// Part 2
type Graph struct {
	Adj      map[int][]int
	Indegree map[int]int
}

func NewGraph(pages []int) *Graph {
	// init a new graph with pages acting as nodes
	g := &Graph{
		Adj:      make(map[int][]int),
		Indegree: make(map[int]int),
	}

	// default set indegree
	for _, page := range pages {
		g.Indegree[page] = 0
	}
	return g
}

func (g *Graph) AddEdge(from, to int) {
	// Adds a directed edge and updates Adj, Indegree states

	g.Adj[from] = append(g.Adj[from], to)
	g.Indegree[to]++
}
