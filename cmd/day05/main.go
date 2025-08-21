package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sotiri-geo/adventOfCode2024/internal/day05"
)

func main() {
	order, errOrder := os.ReadFile("./inputs/day05-order.txt")
	updates, errUpdates := os.ReadFile("./inputs/day05-updates.txt")

	if errOrder != nil {
		panic(errOrder)
	}
	if errUpdates != nil {
		panic(errUpdates)
	}

	predecessor := day05.NewPredecessor(strings.Split(strings.TrimSpace(string(order)), "\n"))
	pageUpdates := day05.ParseUpdates(strings.Split(strings.TrimSpace(string(updates)), "\n"))

	fmt.Println("Part 1:", day05.Part1(predecessor, pageUpdates))
}
