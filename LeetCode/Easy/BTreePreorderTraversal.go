package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Ans struct {
	ans []int
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	a := &Ans{}
	a.solve(root)

	return a.ans
}

func (a *Ans) solve(node *TreeNode) {
	if node == nil {
		return
	}

	a.ans = append(a.ans, node.Val)
	a.solve(node.Left)
	a.solve(node.Right)
}
