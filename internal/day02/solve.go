package day02

import (
	"strconv"
	"strings"
)

const (
	lowerBound = 1
	upperBound = 3
)

func IncreaseWithinBounds(firstValue, secondValue int) bool {

	difference := secondValue - firstValue

	return difference >= lowerBound && difference <= upperBound
}

func IsIncreasing(report []int) bool {
	var increasing = true

	for i := 1; i < len(report); i++ {
		increasing = increasing && IncreaseWithinBounds(report[i-1], report[i])
	}

	return increasing
}

func DecreaseWithinBounds(firstValue, secondValue int) bool {
	return IncreaseWithinBounds(secondValue, firstValue)
}

func IsDecreasing(report []int) bool {
	var decreasing = true

	for i := 1; i < len(report); i++ {
		decreasing = decreasing && DecreaseWithinBounds(report[i-1], report[i])
	}

	return decreasing
}

func removeIndex(i int, array []int) []int {
	newArray := []int{}

	for idx, value := range array {
		if idx == i {
			continue
		}
		newArray = append(newArray, value)
	}
	return newArray
}

func IsDecreasingWithTolerance(report []int) bool {
	return true
}

func IsSafe(report []int) bool {
	return IsIncreasing(report) || IsDecreasing(report)
}

func IsSafeWithTolerance(report []int) bool {
	if IsSafe(report) {
		return true
	}
	// run a check after remove any element once
	for i := range report {
		if IsSafe(removeIndex(i, report)) {
			return true
		}
	}
	return false
}

func toInts(strs []string) []int {
	// good to use make(type, n) when you know the array is of fixed size n
	// no extra allocation of memory as we append to array
	nums := make([]int, len(strs))

	for i, s := range strs {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		nums[i] = n
	}
	return nums
}

// Part 1 solution
func Part1(inputs []string) int {
	var safeCount = 0

	for _, report := range inputs {
		parsedReport := toInts(strings.Fields(report))
		if IsSafe(parsedReport) {
			safeCount++
		}
	}

	return safeCount
}

func Part2(inputs []string) int {
	var safeCount = 0

	for _, report := range inputs {
		parsedReport := toInts(strings.Fields(report))
		if IsSafeWithTolerance(parsedReport) {
			safeCount++
		}
	}

	return safeCount
}
