package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sotiri-geo/adventOfCode2024/internal/day04"
)

func main() {
	data, err := os.ReadFile("./inputs/day04.txt")

	if err != nil {
		panic(err)
	}

	inputs := strings.Split(strings.TrimSpace(string(data)), "\n")
	puzzle := day04.To2DMatrix(inputs)
	fmt.Println("Part 1:", day04.Part1(puzzle))
	fmt.Println("Part 2:", day04.Part2(puzzle))
}
