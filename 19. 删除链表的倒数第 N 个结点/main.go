package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}

	first, second := dummy, dummy
	for i := n; i >= 0; i-- {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}

	next := second.Next
	if next != nil {
		second.Next = next.Next
		next.Next = nil
	}

	return dummy.Next
}
