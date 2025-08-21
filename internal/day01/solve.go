package day01

import (
	"errors"
	"slices"
	"strconv"
	"strings"
)

func Distance(a, b int) int {
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	return abs(a - b)
}

func Sum(distances []int) int {
	var total int

	for _, distance := range distances {
		total += distance
	}

	return total
}

type Row struct {
	Left  int
	Right int
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

func split(ids string) (Row, error) {
	idValues := ToInts(strings.Split(ids, "   "))

	if len(idValues) != 2 {
		return Row{}, errors.New("Can not find two ids in a row.")
	}

	return Row{Left: idValues[0], Right: idValues[1]}, nil
}

func parseIds(inputs []string) ([]int, []int) {
	leftIds := []int{}
	rightIds := []int{}

	for _, ids := range inputs {
		row, err := split(ids)
		if err != nil {
			panic(err)
		}
		leftIds = append(leftIds, row.Left)
		rightIds = append(rightIds, row.Right)
	}
	slices.Sort(leftIds)
	slices.Sort(rightIds)

	return leftIds, rightIds
}

func Part1(inputs []string) int {
	leftIds, rightIds := parseIds(inputs)
	var distances = []int{}

	for i := range len(leftIds) {
		distance := Distance(leftIds[i], rightIds[i])
		distances = append(distances, distance)
	}

	return Sum(distances)
}

// Part 2 solutions

func Counter(ids []int) map[int]int {
	countMap := make(map[int]int)

	for _, id := range ids {
		countMap[id] += 1
	}

	return countMap
}

func SimilarityScore(leftIds []int, rightIds []int) int {
	rightCounter := Counter(rightIds)
	total := 0

	for _, id := range leftIds {
		total += (id * rightCounter[id])
	}

	return total
}

func Part2(inputs []string) int {
	leftIds, rightIds := parseIds(inputs)
	return SimilarityScore(leftIds, rightIds)
}
