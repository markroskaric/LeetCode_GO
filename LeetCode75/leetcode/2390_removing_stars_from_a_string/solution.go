package main

import "fmt"

func main() {
	fmt.Println(removeStars("leet**cod*e"))
}

func removeStars(s string) string {
	stack := []rune{}
	for _, c := range s {
		if c != '*' {
			stack = append(stack, c)
		} else if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}
