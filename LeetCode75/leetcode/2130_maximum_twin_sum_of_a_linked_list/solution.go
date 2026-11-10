package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}
	fmt.Println(pairSum(head))
}

func pairSum(head *ListNode) int {
	slow, fast := head, head
	var prev *ListNode

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		nxt := slow.Next
		slow.Next = prev
		prev = slow
		slow = nxt
	}

	ans := 0
	left, right := prev, slow

	for left != nil {
		if left.Val+right.Val > ans {
			ans = left.Val + right.Val
		}
		left = left.Next
		right = right.Next
	}
	return ans
}
