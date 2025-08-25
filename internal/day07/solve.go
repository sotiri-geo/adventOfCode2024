package day07

func IsCalibrated(calibrations []int, currentTotal int, target int) bool {
	if len(calibrations) == 0 {
		return currentTotal == target
	}

	addition := IsCalibrated(calibrations[1:], calibrations[0]+currentTotal, target)
	multiplication := IsCalibrated(calibrations[1:], calibrations[0]*currentTotal, target)

	return addition || multiplication
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
