package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	crr := head

	for crr != nil {
		for crr.Next != nil && crr.Val == crr.Next.Val {
			crr.Next = crr.Next.Next
		}

		crr = crr.Next
	}

	return head
}
