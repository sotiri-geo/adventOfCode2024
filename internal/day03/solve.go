package day03

import (
	"errors"
	"regexp"
	"strconv"
)

const (
	REGEX_MUL         = `mul\(\d{0,3},\d{0,3}\)`
	REGEX_CONDITIONAL = `do?(n\'t)\(\)`
	ENABLE            = "do()"
	DISABLE           = "don't()"
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

	re := regexp.MustCompile(REGEX_MUL)

	return re.FindAllString(input, -1)
}

func sumExpressions(expressions []string) int {
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

func Part1(input string) int {
	return sumExpressions(ExtractMultiply(input))
}

// Part2 we need to extract the do() and don't()

func ExtractConditionalWithMul(input string) []string {

	re := regexp.MustCompile(REGEX_MUL + "|" + REGEX_CONDITIONAL)

	return re.FindAllString(input, -1)
}

func FilterExpressions(expressions []string) []string {
	keep := []string{}
	var enable = true

	for _, exp := range expressions {
		if exp == ENABLE {
			enable = true
			continue
		}

		if exp == DISABLE {
			enable = false
			continue
		}

		if enable {
			keep = append(keep, exp)
		}
	}

	return keep
}
