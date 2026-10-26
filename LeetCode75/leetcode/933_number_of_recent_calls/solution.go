package main

import "fmt"

func main() {
	a := Constructor()

	fmt.Println(a.Ping(1))
	fmt.Println(a.Ping(100))
	fmt.Println(a.Ping(3001))
	fmt.Println(a.Ping(3002))
}

type RecentCounter struct {
	requests []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.requests = append(this.requests, t)

	for this.requests[0] < t-3000 {
		this.requests = this.requests[1:]
	}

	return len(this.requests)
}
