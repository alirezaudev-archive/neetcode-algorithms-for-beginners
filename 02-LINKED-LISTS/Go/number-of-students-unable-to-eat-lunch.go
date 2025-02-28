package main

func main() {
	println(countStudents([]int{1, 1, 1, 0, 0, 1}, []int{1, 0, 0, 0, 1, 1}))
}

// https://leetcode.com/problems/number-of-students-unable-to-eat-lunch/
func countStudents(students []int, sandwiches []int) int {
	studentsList := fromArray(students)
	sandwichesList := fromArray(sandwiches)
	requeue := 0
	for studentsList.Size > 0 && requeue < studentsList.Size {
		student := studentsList.shift()
		if student.Val == sandwichesList.Head.Val {
			sandwichesList.shift()
			requeue = 0
			continue
		}

		studentsList.append(student.Val)
		requeue++
	}

	return requeue
}

type node struct {
	Val  int
	Next *node
	Prev *node
}

type linkedList struct {
	Head *node
	Tail *node
	Size int
}

func fromArray(arr []int) *linkedList {
	if len(arr) == 0 {
		return nil
	}

	head := &node{Val: arr[0], Prev: nil}
	curr := head
	for _, v := range arr[1:] {
		curr.Next = &node{Val: v, Prev: curr}
		curr = curr.Next
	}

	return &linkedList{
		Head: head,
		Tail: curr,
		Size: len(arr),
	}
}

func (this *linkedList) shift() *node {
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

func (this *linkedList) append(val int) {
	newNode := &node{Val: val, Prev: this.Tail}

	if this.Size == 0 {
		this.Head = newNode
	} else {
		this.Tail.Next = newNode
	}

	this.Tail = newNode
	this.Size++
}
