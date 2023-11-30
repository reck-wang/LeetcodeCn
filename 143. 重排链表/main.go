package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 方法一：线性表展开 + 重组
//func reorderList(head *ListNode) {
//	if head == nil {
//		return
//	}
//
//	var temp []*ListNode
//	for node := head; node != nil; node = node.Next {
//		temp = append(temp, node)
//	}
//
//	i, j := 0, len(temp)-1
//	for i < j {
//		temp[i].Next = temp[j]
//		i++
//		if i == j {
//			break
//		}
//		temp[j].Next = temp[i]
//		j--
//	}
//	temp[i].Next = nil
//}

// 寻找中点 + 原地翻转 + 合并
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid := getMiddle(head)
	lHead := head
	rHead := mid.Next
	mid.Next = nil
	rHead = reverse(rHead)
	merger(lHead, rHead)
}

func getMiddle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func merger(head1, head2 *ListNode) {
	var temp1, temp2 *ListNode
	for temp1 != nil && temp2 != nil {
		temp1 = head1.Next
		temp2 = head2.Next

		head1.Next = head2
		head1 = temp1

		head2.Next = head1
		head2 = temp2
	}
}
