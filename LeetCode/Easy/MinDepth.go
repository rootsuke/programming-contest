package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return solveByDFS(root)
}

func solveByDFS(node *TreeNode) int {
	if node == nil {
		return -1
	}

	if node.Left == nil && node.Right == nil {
		return 1
	}

	ld := solveByDFS(node.Left)
	rd := solveByDFS(node.Right)

	if ld == -1 {
		return rd + 1
	}
	if rd == -1 {
		return ld + 1
	}

	if ld <= rd {
		return ld + 1
	} else {
		return rd + 1
	}
}
