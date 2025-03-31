package main

// https://leetcode.com/problems/design-hashmap/

type Node struct {
	key   int
	value int
	next  *Node
}

type MyHashMap struct {
	set [1000003]*Node
}

func Constructor() MyHashMap {
	return MyHashMap{}
}

func (this *MyHashMap) Put(key, value int) {
	tail, curr := this.getNode(key)
	if curr != nil {
		curr.value = value
		return
	}

	node := &Node{key: key, value: value}
	if tail == nil {
		this.set[this.index(key)] = node
	} else {
		tail.next = node
	}
}

func (this *MyHashMap) Remove(key int) {
	prev, curr := this.getNode(key)
	if curr == nil {
		return
	}

	if prev == nil {
		this.set[this.index(key)] = curr.next
	} else {
		prev.next = curr.next
	}
}

func (this *MyHashMap) Get(key int) int {
	_, curr := this.getNode(key)
	if curr == nil {
		return -1
	}

	return curr.value
}

func (this *MyHashMap) getNode(key int) (*Node, *Node) {
	var prev *Node
	curr := this.set[this.index(key)]
	for curr != nil {
		if curr.key == key {
			return prev, curr
		}
		prev = curr
		curr = curr.next
	}
	return prev, nil
}

func (this *MyHashMap) index(key int) int {
	return key % 1000003
}
