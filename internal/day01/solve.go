package day01

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
