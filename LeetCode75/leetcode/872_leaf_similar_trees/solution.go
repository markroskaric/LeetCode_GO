package main

import (
	"fmt"
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	root2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}}

	fmt.Println(leafSimilar(root1, root2))
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	list1 := make([]int, 0)
	list2 := make([]int, 0)
	recursion(root1, &list1)
	recursion(root2, &list2)
	return reflect.DeepEqual(list1, list2)
}

func recursion(node *TreeNode, leafList *[]int) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		*leafList = append(*leafList, node.Val)
		return
	}
	recursion(node.Left, leafList)
	recursion(node.Right, leafList)
}
