package day07

import "testing"

/*
190: 10 19

  10
+   *
19   190

state variables is an index. We will create a binary tree, where we evaluate the leafs to see
if the value is target. if not returns false
*/

func TestIsCalibrated(t *testing.T) {
	cases := []struct {
		Name         string
		Calibrations []int
		Target       int
		Want         bool
	}{
		{Name: "calibrated equation with single operator +", Calibrations: []int{10, 19}, Target: 190, Want: true},
		{Name: "calibrated equation with multiple operators + *", Calibrations: []int{81, 40, 27}, Target: 3267, Want: true},
		{Name: "cannot calibrate equation", Calibrations: []int{17, 5}, Target: 83, Want: false},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			got := IsCalibrated(tt.Calibrations[1:], tt.Calibrations[0], tt.Target)

			if got != tt.Want {
				t.Errorf("got %v, want %v", got, tt.Want)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	t.Run("total of all possible calibrations", func(t *testing.T) {
		input := map[int][]int{
			190:    []int{10, 19},
			3267:   []int{81, 40, 27},
			83:     []int{17, 5},
			156:    []int{15, 6},
			7290:   []int{6, 8, 6, 15},
			161011: []int{16, 10, 13},
			192:    []int{17, 8, 14},
			21037:  []int{9, 7, 18, 13},
			292:    []int{11, 6, 16, 20},
		}

		got := Part1(input)
		want := 3749

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
