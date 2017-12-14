package length_longest_substring

import (
	"testing"
)

func TestLengthLongestSubstring(t *testing.T) {
	tests := []string {
		"abcabcbb",
		"aaaaaaaa",
		"abcde",
		"",
		"pwwkew",
	}
	results := []int {3, 1, 5, 0, 3}
	for i := 0; i < 5; i++ {
		if ret := LengthLongestSubstring(tests[i]); ret != results[i] {
			t.Fatalf("case %d failed actual: %d, expect: %d\n", i, ret, results[i])
		}
	}
}
