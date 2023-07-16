package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Answer struct {
	ans []int
}

func postorderTraversal(root *TreeNode) []int {
	if root != nil {
		return nil
	}

	a := &Answer{}
	a.solve(root)

	return a.ans
}

func (a *Answer) solve(node *TreeNode) {
	if node == nil {
		return
	}

	a.solve(node.Left)
	a.solve(node.Right)

	a.ans = append(a.ans, node.Val)
}
