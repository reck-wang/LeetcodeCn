package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	dummyHead := &ListNode{}
	sum, carry := 0, 0
	for pre := dummyHead; pre != nil; pre = pre.Next {
		sum = carry
		if l1 != nil {
			sum += l1.Val
		}

		if l2 != nil {
			sum += l2.Val
		}

		if l1 != nil || l2 != nil || sum != 0 {
			pre.Next = &ListNode{Val: sum % 10}
		}

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

		carry = sum / 10
	}

	return dummyHead.Next
}
