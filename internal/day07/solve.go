package day07

import (
	"strconv"
	"strings"
)

func IsCalibrated(calibrations []int, currentTotal int, target int) bool {
	if len(calibrations) == 0 {
		return currentTotal == target
	}

	addition := IsCalibrated(calibrations[1:], calibrations[0]+currentTotal, target)
	multiplication := IsCalibrated(calibrations[1:], calibrations[0]*currentTotal, target)

	return addition || multiplication
}

func IsCalibratedWithConcat(calibrations []int, currentTotal int, target int) bool {
	if len(calibrations) == 0 {
		return currentTotal == target
	}

	addition := IsCalibratedWithConcat(calibrations[1:], calibrations[0]+currentTotal, target)
	multiplication := IsCalibratedWithConcat(calibrations[1:], calibrations[0]*currentTotal, target)
	concatenation := IsCalibratedWithConcat(calibrations[1:], concatOp(currentTotal, calibrations[0]), target)

	return addition || multiplication || concatenation
}

func concatOp(a, b int) int {
	power := 1
	temp := b
	for temp > 0 {
		power *= 10
		temp /= 10
	}
	return a*power + b
}

func ToInts(strs []string) []int {
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

func ParseInput(input []string) map[int][]int {
	output := make(map[int][]int)

	for _, row := range input {
		// EOF has an empty line
		if row == "" {
			continue
		}
		before, after, _ := strings.Cut(row, ": ")
		calibrations := ToInts(strings.Split(strings.TrimSpace(after), " "))
		target, _ := strconv.Atoi(before)
		output[target] = calibrations
	}

	return output
}

func Part1(input map[int][]int) int {
	total := 0

	for target, calibrations := range input {
		if IsCalibrated(calibrations[1:], calibrations[0], target) {
			total += target
		}
	}
	return total
}

// Part 2
func Part2(input map[int][]int) int {
	total := 0

	for target, calibrations := range input {
		if IsCalibratedWithConcat(calibrations[1:], calibrations[0], target) {
			total += target
		}
	}
	return total
}
