package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	i, j, best := 0, len(height)-1, 0
	for i < j {
		curWidth := j - i
		curHeight := min(height[i], height[j])
		best = max(best, curWidth*curHeight)

		if height[i] <= height[j] {
			i++
		} else {
			j--
		}

	}
	return best
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
