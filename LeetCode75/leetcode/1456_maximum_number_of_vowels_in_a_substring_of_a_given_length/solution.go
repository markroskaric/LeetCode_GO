package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxVowels("abciiidef", 3))
}

func maxVowels(s string, k int) int {
	maxSub := 0
	curCount := 0

	vowels := map[string]struct{}{
		"a": {}, "e": {}, "i": {}, "o": {}, "u": {},
	}

	for idx := range s {
		if _, exists := vowels[string(s[idx])]; exists {
			curCount++
		}

		if idx >= k-1 {
			if curCount > maxSub {
				maxSub = curCount
			}

			if _, exists := vowels[string(s[idx-k+1])]; exists {
				curCount--
			}
		}
	}

	return maxSub
}
