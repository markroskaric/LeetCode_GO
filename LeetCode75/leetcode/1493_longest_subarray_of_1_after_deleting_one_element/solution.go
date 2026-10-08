package main

import "fmt"

func main() {
	fmt.Println(longestSubarray([]int{1, 1, 0, 1}))
}

func longestSubarray(nums []int) int {
	k := 1
	maxOne, i := 0, 0
	for index, value := range nums {
		if value == 0 {
			k--
		}
		for k < 0 {
			if nums[i] == 0 {
				k++
			}
			i++

		}
		maxOne = max(index-i, maxOne)
	}
	return maxOne
}
