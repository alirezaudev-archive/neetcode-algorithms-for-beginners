package main

// https://leetcode.com/problems/lru-cache/

type listNode struct {
	k    int
	v    int
	next *listNode
	prev *listNode
}

type LRUCache struct {
	capacity int
	data     map[int]*listNode
	head     *listNode
	tail     *listNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		data:     map[int]*listNode{},
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.data[key]; ok {
		this.moveToTail(node)
		return node.v
	}

	return -1
}

func (this *LRUCache) moveToTail(node *listNode) {
	if node == this.tail {
	    return
	}

    if this.head == node {
        this.head = this.head.next
    } else {
        node.next.prev = node.prev
        node.prev.next = node.next
    }

    node.next = nil
    node.prev = this.tail
    this.tail.next = node
    this.tail = node
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.data[key]; ok {
		node.v = value
		this.moveToTail(node)
		return
	}

	node := &listNode{k: key, v: value, prev: this.tail}
	this.data[key] = node
	length := len(this.data)
	if length == 1 {
		this.head, this.tail = node, node
		return
	}

	this.tail.next = node
	this.tail = node
	if length > this.capacity {
		delete(this.data, this.head.k)
		this.head = this.head.next
		this.head.prev = nil
	}
}
