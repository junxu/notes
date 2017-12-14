package two_sum

import (
	"testing"
)

func TestTwoSumV1(t *testing.T) {
	tests := [][]int{
	    []int{3, 2, 4},
	    []int{0, 8, 7, 3, 3, 4, 2},
	    []int{0,1},
	}
	targets := []int{6, 11, 1}
	results := [][]int{
	    []int{1, 2},
	    []int{1, 3},
	    []int{0, 1},
	}
	caseNum := 3
	for i := 0; i < caseNum; i++ {
		if ret := TwoSumV1(tests[i], targets[i]); ret[0] != results[i][0] &&  ret[1] != results[i][1] {
			t.Fatalf("case %d fais: %v\n", i, ret)
		}
	}
}

func TestD(t *testing.T) {
	tests := [][]int{
	    []int{3, 2, 4},
	    []int{0, 8, 7, 3, 3, 4, 2},
	    []int{0,1},
	}
	targets := []int{6, 11, 1}
	results := [][]int{
	    []int{1, 2},
	    []int{1, 3},
	    []int{0, 1},
	}
	caseNum := 3
	for i := 0; i < caseNum; i++ {
		if ret := TwoSumV1(tests[i], targets[i]); ret[0] != results[i][0] &&  ret[1] != results[i][1] {
			t.Fatalf("case %d fais: %v\n", i, ret)
		}
	}
}

func TestTwoSumV2(t *testing.T) {
	tests := [][]int{
	    []int{3, 2, 4},
	    []int{0, 8, 7, 3, 3, 4, 2},
	    []int{0,1},
	}
	targets := []int{6, 11, 1}
	results := [][]int{
	    []int{1, 2},
	    []int{1, 3},
	    []int{0, 1},
	}
	caseNum := 3
	for i := 0; i < caseNum; i++ {
		if ret := TwoSumV2(tests[i], targets[i]); ret[0] != results[i][0] &&  ret[1] != results[i][1] {
			t.Fatalf("case %d fais: %v\n", i, ret)
		}
	}
}
