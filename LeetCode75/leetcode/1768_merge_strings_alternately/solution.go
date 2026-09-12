package main

import "fmt"

func main() {
	fmt.Println(mergeAlternately("ab", "pqrs"))
}

func mergeAlternately_alter(word1 string, word2 string) string {
	result := ""
	for i := 0; i < len(word1) || i < len(word2); i++ {
		if i < len(word1) {
			result += string(word1[i])
		}
		if i < len(word2) {
			result += string(word2[i])
		}

	}
	return result
}

func mergeAlternately(word1 string, word2 string) string {
	result := ""
	i := 0
	for ; i < len(word1) && i < len(word2); i++ {
		result += string(word1[i]) + string(word2[i])
	}

	if len(word1) > i {
		result += string(word1[i:])
	} else if len(word2) > i {
		result += string(word2[i:])
	}

	return result
}
