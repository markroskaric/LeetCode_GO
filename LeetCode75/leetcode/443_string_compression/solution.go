package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))
}

func compress(chars []byte) int {
	n := len(chars)
	idx := 0
	i := 0
	for i < n {
		ch := chars[i]
		count := 0
		for i < n && ch == chars[i] {
			count++
			i++
		}
		if count == 1 {
			chars[idx] = ch
			idx++
		} else {
			chars[idx] = ch
			idx++
			for _, digit := range []byte(strconv.Itoa(count)) {
				chars[idx] = digit
				idx++
			}
		}
	}

	return idx
}
