package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := &ListNode{Next: head}
	pre, cur := dummyHead, head

	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			temp := cur.Val
			for cur.Next != nil && cur.Next.Val == temp {
				cur = cur.Next
			}
			next := cur.Next
			cur.Next = nil
			pre.Next = next
			cur = next
		} else {
			pre = pre.Next
			cur = cur.Next
		}
	}

	return dummyHead.Next
}
