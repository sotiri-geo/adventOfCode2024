package day02

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
