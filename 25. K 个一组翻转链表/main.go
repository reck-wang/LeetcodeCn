package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var l int
	for temp := head; temp != nil; temp = temp.Next {
		l++
	}

	return reverse(k, l, head)
}

func reverse(k, l int, node *ListNode) *ListNode {
	if l < k {
		return node
	}

	var prev, next *ListNode
	cur := node
	for i := 0; i < k; i++ {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	node.Next = reverse(k, l-k, next)
	return prev
}
