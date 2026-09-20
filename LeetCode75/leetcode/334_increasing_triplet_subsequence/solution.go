package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}))
}

func increasingTriplet(nums []int) bool {
	first := math.MaxInt32
	secound := math.MaxInt32

	for _, num := range nums {
		if num <= first {
			first = num
		} else if num <= secound {
			secound = num
		} else {
			return true
		}
	}
	return false
}
