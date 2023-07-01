package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	left := 0
	right := len(nums) - 1

	return solve(nums, left, right)
}

func solve(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2

	node := &TreeNode{}
	node.Val = nums[mid]
	node.Left = solve(nums, left, mid-1)
	node.Right = solve(nums, mid+1, right)

	return node
}
