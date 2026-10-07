package main

import "fmt"

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
}

func longestOnes(nums []int, k int) int {
	i := 0
	maxwindow := 0
	for index, v := range nums {
		if v == 0 {
			k--
		}
		for k < 0 {

			if nums[i] == 0 {
				k++
			}
			i++

		}
		maxwindow = max(index-i+1, maxwindow)

	}
	return maxwindow
}
