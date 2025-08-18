package day03

import (
	"errors"
	"regexp"
	"strconv"
)

var ErrInsufficientMultiplyArgs = errors.New("cannot multiply without two numbers.")

func Multiply(expression string) (int, error) {
	// We assume that the expression here can be evaluated
	re := regexp.MustCompile(`\d+`)

	matches := re.FindAllString(expression, -1)

	if len(matches) != 2 {
		return 0, ErrInsufficientMultiplyArgs
	}

	nums := make([]int, len(matches))

	for i, value := range matches {
		num, _ := strconv.Atoi(value)
		nums[i] = num
	}

	return nums[0] * nums[1], nil
}

func ExtractMultiply(input string) []string {

	re := regexp.MustCompile(`mul\(\d{0,3},\d{0,3}\)`)

	return re.FindAllString(input, -1)
}

func Part1(input string) int {
	expressions := ExtractMultiply(input)

	var total = 0

	for _, expression := range expressions {
		num, err := Multiply(expression)
		if err != nil {
			panic(err)
		}
		total += num
	}

	return total
}
