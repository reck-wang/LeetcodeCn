package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 方法一 自顶向下归并
//func sortList(head *ListNode) *ListNode {
//	return sort(head, nil)
//}
//
//func sort(head, tail *ListNode) *ListNode {
//	if head == nil {
//		return head
//	}
//
//	if head.Next == tail {
//		head.Next = nil
//		return head
//	}
//
//	fast, slow := head, head
//	for fast != tail {
//		fast= fast.Next
//		if fast != tail {
//			fast = fast.Next
//		}
//		slow = slow.Next
//	}
//
//	return merge(sort(head, slow), sort(slow, tail))
//}
//
//func merge(first, second *ListNode) *ListNode {
//	dummy := &ListNode{Next: nil}
//	pre, temp1, temp2 := dummy, first, second
//	for temp1 != nil && temp2 != nil {
//		if temp1.Val <= temp2.Val {
//			pre.Next = temp1
//			temp1 = temp1.Next
//		} else {
//			pre.Next = temp2
//			temp2 = temp2.Next
//		}
//		pre = pre.Next
//	}
//
//	if temp1 != nil {
//		pre.Next = temp1
//	} else {
//		pre.Next = temp2
//	}
//
//	return dummy.Next
//}

// 方法二 自底向上归并
func sortList(head *ListNode) *ListNode {

}
