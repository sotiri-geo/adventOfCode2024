package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sotiri-geo/adventOfCode2024/internal/day07"
)

func main() {
	data, err := os.ReadFile("./inputs/day07.txt")

	if err != nil {
		panic(err)
	}

	indexByLineData := strings.Split(string(data), "\n")
	parsedInput := day07.ParseInput(indexByLineData)

	fmt.Println("Part 1:", day07.Part1(parsedInput))
	fmt.Println("Part 2:", day07.Part2(parsedInput))

}
