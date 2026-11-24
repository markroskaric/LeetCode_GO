package main

import (
	"fmt"

	tree "LeetCode75/leetcode/tree"
)

func main() {
	root := &tree.TreeNode{
		Val: 3,
		Left: &tree.TreeNode{
			Val:  5,
			Left: &tree.TreeNode{Val: 6},
			Right: &tree.TreeNode{
				Val:   2,
				Left:  &tree.TreeNode{Val: 7},
				Right: &tree.TreeNode{Val: 4},
			},
		},
		Right: &tree.TreeNode{
			Val:   1,
			Left:  &tree.TreeNode{Val: 0},
			Right: &tree.TreeNode{Val: 8},
		},
	}

	p := root.Left  // 5
	q := root.Right // 1
	fmt.Println(lowestCommonAncestor(root, p, q))
}

func lowestCommonAncestor(root, p, q *tree.TreeNode) *tree.TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}
