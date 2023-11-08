package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// recursion
//func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//	if list1 == nil {
//		return list2
//	}
//
//	if list2 == nil {
//		return list1
//	}
//
//	if list1.Val < list2.Val {
//		list1.Next = mergeTwoLists(list1.Next, list2)
//		return list1
//	}
//
//	list2.Next = mergeTwoLists(list1, list2.Next)
//	return list2
//}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var dummyHead = &ListNode{}
	var cur = dummyHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}

	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}

	return dummyHead.Next
}
