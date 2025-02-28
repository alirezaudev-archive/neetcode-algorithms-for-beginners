package linkedlist

import "fmt"

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

func (head *Node) Print() {
	curr := head
	str := " "
	for curr != nil {
		fmt.Print(curr.Val)
		if curr.Next != nil && curr.Next.Prev == curr {
			str += "<"
		}
		if curr.Next != nil {
			fmt.Print(str + "-> ")
		}
		curr = curr.Next
		str = " "
	}
	fmt.Println()
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func FromArray(arr []int) *LinkedList {
	if len(arr) == 0 {
		return nil
	}

	head := &Node{Val: arr[0], Prev: nil}
	curr := head
	for _, v := range arr[1:] {
		curr.Next = &Node{Val: v, Prev: curr}
		curr = curr.Next
	}

	return &LinkedList{
		Head: head,
		Tail: curr,
		Size: len(arr),
	}
}

func (this *LinkedList) Print() {
	curr := this.Head
	fmt.Printf("Size: %d | ", this.Size)
	str := " "
	for curr != nil {
		fmt.Print(curr.Val)
		if curr.Next != nil && curr.Next.Prev == curr {
			str += "<"
		}
		if curr.Next != nil {
			fmt.Print(str + "-> ")
		}
		curr = curr.Next
		str = " "
	}
	fmt.Println()
}

func (this *LinkedList) Reverse() *LinkedList {
	var prev *Node
	curr := this.Head
	for curr != nil {
		next := curr.Next
		curr.Prev = next
		curr.Next = prev
		prev = curr
		curr = next
	}

	this.Tail = this.Head
	this.Head = prev

	return this
}

func (this *LinkedList) Get(index int) int {
	if index >= this.Size || index < 0 {
		return -1
	}
	curr := this.Head
	for index != 0 {
		index--
		curr = curr.Next
	}

	return curr.Val
}

func (this *LinkedList) AddAtHead(val int) {
	newNode := &Node{Val: val, Next: this.Head}

	if this.Size == 0 {
		this.Tail = newNode
	} else {
		this.Head.Prev = newNode
	}

	this.Head = newNode
	this.Size++
}

func (this *LinkedList) Append(val int) {
	newNode := &Node{Val: val, Prev: this.Tail}

	if this.Size == 0 {
		this.Head = newNode
	} else {
		this.Tail.Next = newNode
	}

	this.Tail = newNode
	this.Size++
}

func (this *LinkedList) Insert(index int, val int) {
	if index > this.Size || index < 0 {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
	} else if index == this.Size {
		this.Append(val)
	} else {
		curr := this.Head
		for index > 1 {
			index--
			curr = curr.Next
		}

		newNode := &Node{Val: val, Next: curr.Next, Prev: curr}
		curr.Next.Prev = newNode
		curr.Next = newNode
		this.Size++
	}
}

func (this *LinkedList) Shift() *Node {
	if this.Size == 0 {
		return nil
	}

	tmp := this.Head
	this.Head = this.Head.Next
	if this.Size == 1 {
		this.Tail = this.Head
	}
	this.Size--

	return tmp
}

func (this *LinkedList) Delete(index int) {
	if index >= this.Size || index < 0 {
		return
	}

	if index == 0 {
		this.Head = this.Head.Next
		if this.Size == 1 {
			this.Tail = this.Head
		}
	} else if index+1 == this.Size {
		this.Tail = this.Tail.Prev
		this.Tail.Next = nil
		if this.Size == 1 {
			this.Head = this.Tail
		}
	} else {
		curr := this.Head
		for index > 0 {
			index--
			curr = curr.Next
		}

		curr.Prev.Next = curr.Next
		curr.Next.Prev = curr.Prev
	}

	this.Size--
}
