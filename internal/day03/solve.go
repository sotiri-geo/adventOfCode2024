package day03

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

const (
	REGEX_MUL         = `mul\(\d{0,3},\d{0,3}\)`
	REGEX_CONDITIONAL = `do(n\'t){0,1}\(\)`
	ENABLE            = "do()"
	DISABLE           = "don't()"
)

// Compile regex patterns once at package init for better performance
var (
	mulRegex                = regexp.MustCompile(REGEX_MUL)
	conditionalWithMulRegex = regexp.MustCompile(REGEX_MUL + "|" + REGEX_CONDITIONAL)
	numberRegex             = regexp.MustCompile(`\d+`)
)

var ErrInsufficientMultiplyArgs = errors.New("cannot multiply without two numbers.")

func Multiply(expression string) (int, error) {
	// We assume that the expression here can be evaluated
	matches := numberRegex.FindAllString(expression, -1)

	if len(matches) != 2 {
		return 0, ErrInsufficientMultiplyArgs
	}

	// preallocate with known capacity
	nums := make([]int, 0, 2)

	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			return 0, fmt.Errorf("invalid number %q: %w", match, err)
		}
		nums = append(nums, num)
	}

	return nums[0] * nums[1], nil
}

func ExtractMultiply(input string) []string {
	return mulRegex.FindAllString(input, -1)
}

func sumExpressions(expressions []string) (int, error) {
	total := 0

	for _, expression := range expressions {
		num, err := Multiply(expression)
		if err != nil {
			return 0, fmt.Errorf("failed to multiply expression %q: %w", expression, err)
		}
		total += num
	}

	return total, nil
}

func Part1(input string) int {
	total, _ := sumExpressions(ExtractMultiply(input))
	return total
}

// Part2 we need to extract the do() and don't()

func ExtractConditionalWithMul(input string) []string {
	return conditionalWithMulRegex.FindAllString(input, -1)
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

func Part2(input string) int {
	total, _ := sumExpressions(FilterExpressions(ExtractConditionalWithMul(input)))
	return total
}
