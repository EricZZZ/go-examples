package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func calculateDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(calculateDepth(root.Left), calculateDepth(root.Right)) + 1
}
