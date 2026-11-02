package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 7, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6}}}}}}}
	fmt.Println(deleteMiddle(&list))
}

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	slow := head
	fast := head
	var prev *ListNode = nil

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = slow.Next

	return head
}
