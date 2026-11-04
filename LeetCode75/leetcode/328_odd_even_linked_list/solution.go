package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}
	fmt.Println(oddEvenList(head))
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	firstEven := head.Next
	oddWalker := head
	evenWalker := head.Next

	for {
		if evenWalker == nil || evenWalker.Next == nil {
			oddWalker.Next = firstEven
			break
		}
		oddWalker.Next = oddWalker.Next.Next
		evenWalker.Next = evenWalker.Next.Next
		oddWalker = oddWalker.Next
		evenWalker = evenWalker.Next
	}
	return head
}
