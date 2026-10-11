package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}

func pivotIndex(nums []int) int {
	totalSum, leftSum := 0, 0

	for _, value := range nums {
		totalSum += value
	}

	for index, value := range nums {
		totalSum = totalSum - value
		if totalSum == leftSum {
			return index
		}
		leftSum += value
	}
	return -1
}
