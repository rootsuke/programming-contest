package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			current = current.Next
			list1 = list1.Next
		} else {
			current.Next = list2
			current = current.Next
			list2 = list2.Next
		}
	}

	if list1 == nil {
		current.Next = list2
	} else {
		current.Next = list1
	}

	return head.Next
}

func main() {
	a2 := &ListNode{Val: 3, Next: nil}
	a1 := &ListNode{Val: 1, Next: a2}

	b2 := &ListNode{Val: 4, Next: nil}
	b1 := &ListNode{Val: 2, Next: b2}

	node := mergeTwoLists(a1, b1)

	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}
