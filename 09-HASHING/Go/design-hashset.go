package main

// https://leetcode.com/problems/design-hashset/

type Node struct {
	key  int
	next *Node
}

type MyHashSet struct {
	set [10007]*Node
}

func Constructor() MyHashSet {
	return MyHashSet{}
}

func (this *MyHashSet) Add(key int) {
	tail, curr := this.get(key)
	if curr != nil {
		return
	}
	node := &Node{key: key}
	if tail == nil {
		this.set[this.index(key)] = node
	} else {
		tail.next = node
	}
}

func (this *MyHashSet) Remove(key int) {
	prev, curr := this.get(key)
	if curr == nil {
		return
	}

	if prev == nil {
		this.set[this.index(key)] = curr.next
	} else {
		prev.next = curr.next
	}
}

func (this *MyHashSet) Contains(key int) bool {
	_, curr := this.get(key)
	return curr != nil
}

func (this *MyHashSet) get(key int) (*Node, *Node) {
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

func (this *MyHashSet) index(key int) int {
	return key % 10007
}
