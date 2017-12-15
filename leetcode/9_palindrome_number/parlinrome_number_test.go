package palindrome_number

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []int{-121, -132, -1, -111, -1111, -1221, -12324323}
	results := []bool{true, false, true, true, true, true, false}
	for i, v := range tests {
		if ret :=IsPalindrome(v); ret != results[i] {
			t.Fatalf("case %d failed, actaul is %v, expect is %v\n", tests[i], ret, results[i])
		}
	}
}

func TestIsPalindrome2(t *testing.T) {
	tests := []int{-121, -132, -1, -111, -1111, -1221, -12324323}
	results := []bool{true, false, true, true, true, true, false}
	for i, v := range tests {
		if ret :=IsPalindrome2(v); ret != results[i] {
			t.Fatalf("case %d failed, actaul is %v, expect is %v\n", tests[i], ret, results[i])
		}
	}
}
