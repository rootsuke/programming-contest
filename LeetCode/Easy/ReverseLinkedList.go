package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	for head != nil {
		next := head.Next
		head.Next = newHead
		newHead = head
		head = next
	}

	return newHead

	// return reverse(head, nil)
}

// func reverse(head, newHead *ListNode) *ListNode {
// 	if head == nil {
// 		return newHead
// 	}

// 	next := head.Next
// 	head.Next = newHead

// 	return reverse(next, head)
// }
