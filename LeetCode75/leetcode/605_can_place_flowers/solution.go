package main

import "fmt"

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))

	//fmt.Println(len([]int{1, 1, 1, 1, 1}))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	lena := len(flowerbed)
	for i := 0; i < lena; i++ {

		left := i == 0 || flowerbed[i-1] == 0
		right := i == lena-1 || flowerbed[i+1] == 0

		if left && right && flowerbed[i] == 0 {
			flowerbed[i] = 1
			n--
		}
	}
	return n <= 0
}
