package main

import "fmt"

func main() {
	fmt.Println(reverseVowels("IceCreAm"))
}

func reverseVowels(s string) string {
	left, right := 0, len(s)-1
	runes := []rune(s)
	for left < right {
		if containsViwels(runes[left]) && containsViwels(runes[right]) {
			runes[left], runes[right] = runes[right], runes[left]
			left++
			right--
		}
		if !containsViwels(runes[left]) {
			left++
		}
		if !containsViwels(runes[right]) {
			right--
		}
	}

	return string(runes)
}

func containsViwels(s rune) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

	for _, element := range vowels {
		if element == s {
			return true
		}
	}
	return false
}
