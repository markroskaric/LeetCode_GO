package main

import "fmt"

func main() {
	fmt.Println(findDifference([]int{1, 2, 3}, []int{2, 4, 6}))
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	set1 := map[int]bool{}
	set2 := map[int]bool{}

	for _, n := range nums1 {
		set1[n] = true
	}
	for _, n := range nums2 {
		set2[n] = true
	}
	diff1, diff2 := []int{}, []int{}

	for k := range set1 {
		if !set2[k] {
			diff1 = append(diff1, k)
		}
	}
	for k := range set2 {
		if !set1[k] {
			diff2 = append(diff2, k)
		}
	}
	return [][]int{diff1, diff2}
}
