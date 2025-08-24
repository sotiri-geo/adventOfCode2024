package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sotiri-geo/adventOfCode2024/internal/day06"
)

func main() {
	data, err := os.ReadFile("./inputs/day06.txt")

	if err != nil {
		panic(err)
	}

	inputMap := day06.To2DMatrix(strings.Split(strings.TrimSpace(string(data)), "\n"))

	fmt.Println("Part 1:", day06.Part1(inputMap))
}
