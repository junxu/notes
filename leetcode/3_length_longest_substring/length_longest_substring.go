package length_longest_substring

func LengthLongestSubstring(s string) int {
	left, right, max := 0, 0, 0
	temp := make(map[byte]int)

	for right < len(s) {
		if r, ok := temp[s[right]]; ok && left <= r {
			left = temp[s[right]] + 1
		} else {
			if right -left + 1 > max {
				max = right - left + 1
			}
		}
		
		temp[s[right]] = right
		right = right + 1
	}

	return max
}
