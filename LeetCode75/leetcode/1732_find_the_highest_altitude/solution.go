package main

import "fmt"

func main() {
	fmt.Println(largestAltitude([]int{-5, 1, 5, 0, -7}))
}

func largestAltitude(gain []int) int {
	sum, maxAltitude := 0, 0
	for _, value := range gain {
		sum = sum + value
		maxAltitude = max(maxAltitude, sum)
	}
	return maxAltitude
}
