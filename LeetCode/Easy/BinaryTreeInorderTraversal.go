// import "fmt"

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := []int{}

	ans = solve(root, ans)

	return ans
}

func solve(node *TreeNode, ans []int) []int {

	// fmt.Printf("val: %v\n", node.Val)
	// fmt.Printf("left: %v\n", node.Left)
	// fmt.Printf("right: %v\n", node.Right)
	// fmt.Println()

	if node.Left != nil {
		ans = solve(node.Left, ans)
	}

	ans = append(ans, node.Val)

	if node.Right != nil {
		ans = solve(node.Right, ans)
	}

	return ans
}
