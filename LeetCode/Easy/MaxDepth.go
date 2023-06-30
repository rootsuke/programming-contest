package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return solve(root)
}

func solve(node *TreeNode) int {
	if node == nil {
		return 0
	}

	ld := solve(node.Left)
	rd := solve(node.Right)

	if ld >= rd {
		return ld + 1
	} else {
		return rd + 1
	}
}
