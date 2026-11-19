package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 10, Right: &TreeNode{Val: -3, Right: &TreeNode{Val: 11}}, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: -2}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 1}}}}
	fmt.Println(pathSum(root, 8))
}

func pathSum(root *TreeNode, targetSum int) int {
	return countPaths(root, targetSum, []int{})
}

func countPaths(root *TreeNode, targetValue int, path []int) int {
	counter := 0
	if root == nil {
		return counter
	}

	path = append(path, root.Val)
	sum := 0

	for i := len(path) - 1; i >= 0; i-- {
		sum += path[i]
		if sum == targetValue {
			counter++
		}
	}

	return counter + countPaths(root.Left, targetValue, path) + countPaths(root.Right, targetValue, path)
}
