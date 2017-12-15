package palindrome_number

import "strconv"


func IsPalindrome(x int) bool {
	if x < 0 {
		x = -x
	}
	s := strconv.Itoa(x)
	i, j := 0, 0
	if len(s) % 2 == 1 {
		i, j = len(s)/2, len(s) / 2
	} else {
		i, j = len(s)/2 - 1, len(s) /2
	}

	for i >=0 && j <= len(s)-1 {
		if s[i] != s[j] {
			return false
		}
		i--
		j++
	}
	return true
}

func IsPalindrome2(x int) bool {
	if x < 0 {
		x = -x
	}

	if x == 0 {
		return true
	}

	temp :=  x
	res := 0
	for temp != 0 {
		res = res * 10 + temp % 10
		temp = temp / 10
	}

	return res == x
}
