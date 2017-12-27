package container_water

func MaxArea(height []int) int {
	max, l, r := 0, 0, len(height) - 1

	for l < r {
		temp := 0
		if height[l] >= height[r] {
			temp = height[r] * (r-l)
			r--
		} else {
			temp = height[l] * (r - l)
			l++
		}
		if temp > max {
			max = temp
		}
	}
	return max
}
