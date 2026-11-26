package main

import (
	"fmt"

	tree "LeetCode75/leetcode/tree"
)

func main() {
	root := &tree.TreeNode{
		Val: 1,
		Left: &tree.TreeNode{
			Val:   2,
			Right: &tree.TreeNode{Val: 5},
		},
		Right: &tree.TreeNode{
			Val:   3,
			Right: &tree.TreeNode{Val: 4},
		},
	}

	fmt.Println(rightSideView(root))
}

func rightSideView(root *tree.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := []int{}
	queue := []*tree.TreeNode{root}

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == n-1 {
				ans = append(ans, node.Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}
	}
	return ans
}
