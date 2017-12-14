package two_sum

import (
	"fmt"
)

func TwoSumV1(nums []int, target int) []int {
	for i:=0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if sum == target { 
				return []int{i, j}
			}
		}
	}
	return nil
}


func TwoSumV2(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, key := range nums {
		fmt.Println(i, key)
		if j, ok := hashMap[target - key]; ok {
			return []int{j, i}
		}
		hashMap[key] = i
	}
	return nil
}
