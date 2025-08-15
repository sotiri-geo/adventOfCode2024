package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sotiri-geo/adventOfCode2024/internal/day02"
)

func main() {
	data, err := os.ReadFile("./inputs/day02.txt")

	if err != nil {
		panic(err)
	}

	inputs := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Println("Part 1:", day02.Part1(inputs))
}
