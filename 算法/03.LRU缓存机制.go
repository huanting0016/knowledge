package main

type LinkNode struct {
	key   int
	value int
	pre   *LinkNode
	next  *LinkNode
}

type LRUCache struct {
	m    map[int]*LinkNode
	cap  int
	head *LinkNode
	tail *LinkNode
}

func Constructor(cap int) LRUCache {
	head := &LinkNode{-1, -1, nil, nil}
	tail := &LinkNode{-1, -1, nil, nil}
	head.next = tail
	tail.pre = head
	cache := LRUCache{make(map[int]*LinkNode), cap, head, tail}
	return cache

}

func (this *LRUCache) AddNode(node *LinkNode) {
	node.pre = this.head
	node.next = this.head.next
	this.head.next = node
	node.next.pre = node
}

func (this *LRUCache) RemoveNode(node *LinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) MoveToHead(node *LinkNode) {
	this.RemoveNode(node)
	this.AddNode(node)
}

func (this *LRUCache) Get(key int) int {
	m := this.m
	if node, ok := m[key]; ok {
		this.MoveToHead(node)
		return node.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	m := this.m
	if node, ok := m[key]; ok {
		node.value = value
		this.MoveToHead(node)
	} else {
		n := &LinkNode{key, value, nil, nil}
		if len(m) >= this.cap {
			delete(m, this.tail.pre.key)
			this.RemoveNode(this.tail.pre)
		}
		m[key] = n
		this.AddNode(n)
	}
}
