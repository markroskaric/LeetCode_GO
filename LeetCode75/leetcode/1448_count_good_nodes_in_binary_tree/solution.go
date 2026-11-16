package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 2}}, Right: nil}
	fmt.Println(goodNodes(root))
}

func goodNodes(root *TreeNode) int {
	return isGoodNode(root, root.Val)
}

func isGoodNode(node *TreeNode, maxVal int) int {
	if node == nil {
		return 0
	}
	count := 0

	if node.Val >= maxVal {
		count = 1
		maxVal = node.Val
	}

	count += isGoodNode(node.Left, maxVal)
	count += isGoodNode(node.Right, maxVal)

	return count
}
