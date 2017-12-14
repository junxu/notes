package longest_palindrome

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []string {
		"babad",
		"cbbd",
		"bbbbbbb",
		"abababab",
		"abcfcdfgfdc",
	}

	results := []string {
		"bab",
		"bb",
		"bbbbbbb",
		"abababa",
		"cdfgfdc",
	}

	for i := 0; i < 5; i++ {
		if ret := LongestPalindrome(tests[i]); ret != results[i] {
			t.Fatalf("case %d failed, actual: %v, expect: %v\n", i, ret, results[i])
		}
	}
}
