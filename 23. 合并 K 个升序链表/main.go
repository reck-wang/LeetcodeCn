package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	pq := &priorityQueue{nodes: []*ListNode{}}
	for _, list := range lists {
		if list != nil {
			heap.Push(pq, list)
		}
	}

	dummyHead := &ListNode{}
	cur := dummyHead
	for pq.Len() > 0 {
		node := heap.Pop(pq).(*ListNode)
		cur.Next = node

		if node.Next != nil {
			heap.Push(pq, node.Next)
		}
		cur = cur.Next
	}

	return dummyHead.Next
}

type priorityQueue struct {
	nodes []*ListNode
}

func (p *priorityQueue) Len() int {
	return len(p.nodes)
}

func (p *priorityQueue) Less(i, j int) bool {
	return p.nodes[i].Val < p.nodes[j].Val
}

func (p *priorityQueue) Swap(i, j int) {
	p.nodes[i], p.nodes[j] = p.nodes[j], p.nodes[i]
}

func (p *priorityQueue) Push(i interface{}) {
	p.nodes = append(p.nodes, i.(*ListNode))
}

func (p *priorityQueue) Pop() interface{} {
	temp := len(p.nodes) - 1
	res := p.nodes[temp]
	p.nodes = p.nodes[:temp]
	return res
}
