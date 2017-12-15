package median_two_sorted_array

import "testing"

func TestMedianTwoSortedArray(t *testing.T) {
	test_list1 := [][]int {
		[]int{1, 3},
		[]int{1, 2},
		[]int{1, 2, 3, 4, 5, 6},
	}

	test_list2 := [][]int {
		[]int{2},
		[]int{3, 4},
		[]int{},
	}

	results := []float32{2.0, 2.5, 3.5}

	for i := 0; i < len(test_list1); i++ {
		ret := MedianTwoSortedArray(test_list1[i], test_list2[i])
		if ret != results[i] {
			t.Fatalf("case %d failed, actual is %v, expect is %v\n", i, ret, results[i])
		}
	}
}


func TestFindIndexTwoSortedArray(t *testing.T) {
	test_list1 := [][]int {
		[]int{1, 3},
		[]int{1, 2},
		[]int{1,2,3,4,5,6},
		[]int{1,2,3,4,5,6},
	}

	test_list2 := [][]int {
		[]int{2},
		[]int{3, 4},
		[]int{7,8,9,10},
		[]int{7,8,9,10},
	}

	poss := []int{2, 4, 6, 10}

	results := [][]int{
		[]int{2, 0},
		[]int{2, 1},
		[]int{1, 4},
		[]int{2, 3},
	}

	for i := 0; i < 4; i++ {
		index, pos := FindIndexTwoSortedArray(test_list1[i], test_list2[i], poss[i])
		if index != results[i][0] || pos != results[i][1] {
			t.Fatalf("case %d failed, actaul is %d-%d, expect %d-%d\n", i, index, pos, results[i][0], results[i][1])
		}
	}
}


func TestFindPos(t *testing.T) {
	l := []int{1, 4, 20 , 21, 50}
	tests := []int{-1, 60, 20, 50, 17}
	results := []int{0, 5, 3, 5, 2}

	for i, v := range(tests) {
		if ret :=findPos(l, 0, len(l)-1, v); ret != results[i] {
			t.Fatalf("case %d failed, expect %d, actual %d\n", i, results[i], ret)
		}
	}
	if ret :=findPos(l, 1, len(l)-1, 2); ret != 1 {
		t.Fatalf("case %d failed, expect %d, actual %d\n", 5, 1, ret)
	}
	if ret :=findPos(l, 1, len(l)-1, -2); ret != 1 {
		t.Fatalf("case %d failed, expect %d, actual %d\n", 6, 1, ret)
	}
}

func TestFindPos2(t *testing.T) {
	l := []int{1, 4, 20 , 21, 50}
	tests := []int{-1, 60, 20, 50, 17}
	results := []int{0, 5, 3, 5, 2}

	for i, v := range(tests) {
		ret :=findPos2(l, v)
		if ret != results[i] {
			t.Fatalf("case %d failed, expect %d, actual %d\n", i, results[i], ret)
		}
	}
}
