package main

import "fmt"

type Solution struct{}

func main() {
	sol := &Solution{}
	fmt.Println(sol.CloseStrings("abc", "bca"))
}

func (s *Solution) CloseStrings(word1 string, word2 string) bool {
	length := len(word1)
	if length != len(word2) {
		return false
	}

	chars1 := make([]int, 26)
	chars2 := make([]int, 26)
	values1 := make([]int, length+1)
	values2 := make([]int, length+1)

	s.process(chars1, values1, word1)
	s.process(chars2, values2, word2)
	fmt.Println(chars1)
	fmt.Println(values1)
	for i := 0; i < 26; i++ {
		if chars1[i] > 0 && chars2[i] == 0 {
			return false
		}
	}
	for i := 0; i < length; i++ {
		if values1[i] != values2[i] {
			return false
		}
	}

	return true
}

func (s *Solution) process(chars []int, values []int, word string) {
	for _, c := range word {
		chars[c-'a']++
		values[chars[c-'a']]++
	}
}
