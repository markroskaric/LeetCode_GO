package main

import "fmt"

func main() {
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	biggest := candies[0]

	for _, candie := range candies {
		if biggest < candie {
			biggest = candie
		}
	}

	number := biggest - extraCandies
	for i, kidCandies := range candies {
		if kidCandies < number {
			result[i] = false
		} else {
			result[i] = true
		}
	}
	return result
}
