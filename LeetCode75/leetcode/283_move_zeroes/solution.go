package main

import "fmt"

func main() {
	nums := []int{1, 0, 1}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	left := 0

	for right := 0; right < len(nums); right++ {
		for left < right && nums[left] != 0 {
			left++
		}

		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
}
