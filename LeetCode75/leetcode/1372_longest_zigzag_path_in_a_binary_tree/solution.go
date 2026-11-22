package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1, Right: &TreeNode{Val: 1}, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 1, Right: &TreeNode{Val: 1}, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 1}}}}}
	fmt.Println(longestZigZag(root))
}

func longestZigZag(root *TreeNode) int {
	return max(dfs(root.Left, true, 0), dfs(root.Right, false, 0))
}

func dfs(node *TreeNode, isLeft bool, count int) int {
	if node == nil {
		return count
	}

	if isLeft {
		return max(dfs(node.Right, false, count+1), dfs(node.Left, true, 0))
	}

	return max(dfs(node.Right, false, 0), dfs(node.Left, true, count+1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
