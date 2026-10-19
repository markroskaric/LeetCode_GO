package main

import "fmt"

func main() {
	fmt.Println(asteroidCollision([]int{5, 10, -5}))
}

func asteroidCollision(asteroids []int) []int {
	prev := -1
	for i := 0; i < len(asteroids); i++ {
		notSame := true
		if asteroids[i] > 0 {
			prev++
			asteroids[prev] = asteroids[i]
		} else {
			for prev >= 0 && asteroids[prev] > 0 {
				if asteroids[prev] < -asteroids[i] {
					prev--
				} else if asteroids[prev] == -asteroids[i] {
					prev--
					notSame = false
					break

				} else {
					break
				}
			}
		}
		if notSame && (prev < 0 || asteroids[prev] < 0) {
			prev++
			asteroids[prev] = asteroids[i]
		}
	}
	return asteroids[:prev+1]
}
