package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return solve(root.Left, root.Right)
}

func solve(left *TreeNode, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}

	if left.Val != right.Val {
		return false
	}

	return solve(left.Left, right.Right) && solve(left.Right, right.Left)
}
