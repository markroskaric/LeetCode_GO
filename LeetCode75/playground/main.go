package main

import "fmt"

func main() {
	yes := printAll()
	yes("mark")
	yes("je")
	fmt.Println(yes("cool"))

	squareFunc := selfMath(multiply)
	doubleFunc := selfMath(add)

	fmt.Println(squareFunc(4))
	fmt.Println(doubleFunc(4))
	for {
		go fmt.Println(yes("cool"))
	}
}

func printAll() func(string) string {
	text := ""
	return func(word string) string {
		text += word + " "
		return text
	}
}

func multiply(x, y int) int {
	return x * y
}

func add(x, y int) int {
	return x + y
}

func selfMath(mathfunc func(x, y int) int) func(int) int {
	return func(x int) int {
		return mathfunc(x, x)
	}
}
