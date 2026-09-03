package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res, ks int

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	if ks == 0 {
		return
	}
	ks--
	if ks == 0 {
		res = root.Val
	}
	dfs(root.Right)

}

func kthSmallest(root *TreeNode, k int) int {
	ks = k
	dfs(root)
	return res
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 2}
	res := kthSmallest(root, 1)
	println(res)
}
