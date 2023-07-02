package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return solve(root, 0) != -1
}

func solve(node *TreeNode, height int) int {
	if node == nil {
		return height
	}

	ld := solve(node.Left, height+1)
	if ld == -1 {
		return -1
	}
	rd := solve(node.Right, height+1)
	if rd == -1 {
		return -1
	}

	depthGap := 0
	if ld <= rd {
		depthGap = rd - ld
	} else {
		depthGap = ld - rd
	}
	if depthGap > 1 {
		return -1
	}

	if ld <= rd {
		return rd
	} else {
		return ld
	}
}
