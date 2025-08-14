package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sotiri-geo/adventOfCode2024/internal/day01"
)

func main() {
	data, err := os.ReadFile("./inputs/day01.txt")
	if err != nil {
		panic(err)
	}

	input := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Println("Part 1:", day01.Part1(input))
}
