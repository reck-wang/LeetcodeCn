package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	time1, time2 := 1, 1

	for curA != curB {
		curA = curA.Next
		curB = curB.Next
		if curA == nil && time1 == 1 {
			curA = headB
			time1--
		}

		if curB == nil && time2 == 1 {
			curB = headA
			time2--
		}
	}

	return curA
}
