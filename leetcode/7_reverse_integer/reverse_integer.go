package reverse_integer

import "strconv"

func ReverseInteger(x int) int {
	s := strconv.Itoa(x)
	n := 0 
	runes := make([]rune, len(s))
	for _, r := range s {
		runes[n] = r
		n++ 
	}

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	s = string(runes)
	ret, _ := strconv.Atoi(s)

	return ret
}
