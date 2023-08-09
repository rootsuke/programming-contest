package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	solve(root)

	return root
}

func solve(node *TreeNode) {
	if node.Left == nil && node.Right == nil {
		return
	}

	if node.Left != nil {
		solve(node.Left)
	}

	if node.Right != nil {
		solve(node.Right)
	}

	node.Left, node.Right = node.Right, node.Left
}
