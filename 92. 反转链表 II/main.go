package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	pre := dummyHead
	for i := left; i > 1; i-- {
		pre = pre.Next
	}

	cur := pre.Next
	var next *ListNode
	for i := 0; i < right-left; i++ {
		next = cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}

	return dummyHead.Next
}
