package main

type Node struct {
	Key, Val   int
	Prev, Next *Node
}

type LRUCache struct {
	Cap        int
	Cache      map[int]*Node
	Head, Tail *Node
}

func Constructor(capacity int) LRUCache {
	heap := &Node{0, 0, nil, nil}
	tail := &Node{0, 0, nil, nil}
	heap.Next = tail
	tail.Prev = heap

	return LRUCache{
		Cap:   capacity,
		Cache: make(map[int]*Node, capacity+1),
		Head:  heap,
		Tail:  tail,
	}
}

func (l *LRUCache) Get(key int) int {
	if node, has := l.Cache[key]; has {
		l.unLink(node)
		l.appendHead(node)
		return node.Val
	}
	return -1
}

func (l *LRUCache) Put(key, val int) {
	if node, has := l.Cache[key]; has {
		l.unLink(node)
	}

	temp := &Node{key, val, nil, nil}
	l.appendHead(temp)
	l.Cache[key] = temp

	if len(l.Cache) > l.Cap {
		tail := l.removeTail()
		delete(l.Cache, tail.Key)
	}
}

func (l *LRUCache) unLink(node *Node) {
	prev := node.Prev
	next := node.Next

	prev.Next = next
	next.Prev = prev

	node.Prev = nil
	node.Next = nil
}

func (l *LRUCache) appendHead(node *Node) {
	next := l.Head.Next

	node.Prev = l.Head
	node.Next = next

	next.Prev = node
	l.Head.Next = node
}

func (l *LRUCache) removeTail() *Node {
	node := l.Tail.Prev
	prev := node.Prev

	prev.Next = l.Tail
	l.Tail.Prev = prev

	node.Prev = nil
	node.Next = nil

	return node
}
