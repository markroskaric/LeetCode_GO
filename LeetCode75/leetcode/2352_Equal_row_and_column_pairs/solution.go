package main

import "fmt"

func main() {
	fmt.Println(equalPairs([][]int{
		{3, 2, 2},
		{1, 7, 7},
		{2, 7, 7},
	}))
}

func equalPairs(grid [][]int) int {
	n := len(grid)
	m := make(map[[200]int]int)
	arr := [200]int{}
	for i := range n {
		copy(arr[:], grid[i])
		m[arr]++
	}
	res := 0
	for i := range n {
		arr = [200]int{}
		for j := range n {
			arr[j] = grid[j][i]
		}
		if v, ok := m[arr]; ok {
			res += v
		}

	}
	return res
}
