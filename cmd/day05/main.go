package main

import (
	"fmt"
	"os"
	"strings"
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

	parsedOrder := strings.Split(string(order), "\n")
	parsedUpdates := strings.Split(string(updates), "\n")

	fmt.Println(parsedOrder)
	fmt.Println(parsedUpdates)
}
