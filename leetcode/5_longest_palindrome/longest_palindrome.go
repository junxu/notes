package longest_palindrome

func LongestPalindrome(s string) string {
	maxP := ""
	for i := 0; i < len(s); i++ {
		ret1 := extendCenter(s, i, i)
		if len(maxP) < len(ret1) {
			maxP = ret1
		}

		ret2 := extendCenter(s, i, i+1)
		if len(maxP) < len(ret2) {
			maxP = ret2
		}
	}
	return maxP
}

func extendCenter(s string, left, right int) string {
	for left >=0 && right < len(s) && s[left] == s[right] {
		left = left - 1
		right = right + 1
	}
	return s[left+1 : right]
}
