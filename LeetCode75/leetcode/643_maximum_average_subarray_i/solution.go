package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}

func findMaxAverage(nums []int, k int) float64 {
	left, right := 0, 0
	var sum float64 = 0
	maxAvg := math.Inf(-1)

	for right < len(nums) {
		window := right - left + 1
		sum += float64(nums[right])

		if window == k {
			maxAvg = max(sum/float64(k), maxAvg)
			sum -= float64(nums[left])
			left++
		}
		right++
	}
	return maxAvg
}
