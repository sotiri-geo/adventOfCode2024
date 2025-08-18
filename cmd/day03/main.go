package main

import (
	"fmt"
	"os"

	"github.com/sotiri-geo/adventOfCode2024/internal/day03"
)

func main() {
	data, err := os.ReadFile("./inputs/day03.txt")

	if err != nil {
		panic(err)
	}

	input := string(data)

	fmt.Println("Part 1:", day03.Part1(input))
	fmt.Println("Part 2:", day03.Part2(input))
}
