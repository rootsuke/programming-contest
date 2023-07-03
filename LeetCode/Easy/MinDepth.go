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

type Queue struct {
	queue []*TreeNode
}

func (q *Queue) Enqueue(value *TreeNode) {
	q.queue = append(q.queue, value)
}

func (q *Queue) Dequeue() *TreeNode {
	value := q.queue[0]
	q.queue = q.queue[1:]

	return value
}

func (q *Queue) isExists() bool {
	return len(q.queue) != 0
}

func solveByBFS(node *TreeNode) int {
	queue := &Queue{}
	queue.Enqueue(node)

	depth := 1
	for queue.isExists() {
		size := len(queue.queue)
		for i := 0; i < size; i++ {
			crr := queue.Dequeue()

			if crr.Left == nil && crr.Right == nil {
				return depth
			}

			if crr.Left != nil {
				queue.Enqueue(crr.Left)
			}
			if crr.Right != nil {
				queue.Enqueue(crr.Right)
			}
		}

		depth++
	}

	return -1
}
