package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 方法一: 递归
//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//
//	newHead := reverseList(head.Next)
//	head.Next.Next = head
//	head.Next = nil
//
//	return newHead
//}

// 方法二: 迭代
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var dummyhead *ListNode
	pre, temp := dummyhead, head

	for temp != nil && temp.Next != nil {
		next := temp.Next
		temp.Next = pre
		pre = temp
		temp = next
	}

	return pre.Next
}
