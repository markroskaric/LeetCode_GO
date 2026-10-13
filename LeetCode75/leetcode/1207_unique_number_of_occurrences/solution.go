package main

import "fmt"

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2}))
}

func uniqueOccurrences(arr []int) bool {
	set := map[int]int{}

	for _, v := range arr {
		set[v] += 1
	}
	seenValues := make(map[int]int)
	for v, r := range set {
		if _, exists := seenValues[r]; exists {
			return false
		}
		seenValues[r] = v
	}
	return true
}
