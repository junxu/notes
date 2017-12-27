package container_water

import "testing"

func TestMaxArea(t *testing.T) {
	tests := [][]int {
		[]int{1, 1},
		[]int{1, 2},
		[]int{1, 2, 4, 3},
	}
	results := []int{1, 1, 4}

	for i := 0; i < len(tests); i++ {
		if ret := MaxArea(tests[i]); ret != results[i] {
			t.Fatalf("case %d failed, actual is %d, expect is %d\n", i, ret, results[i])
		}
	}
}
