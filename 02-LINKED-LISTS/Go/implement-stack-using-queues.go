package main

// https://leetcode.com/problems/implement-stack-using-queues/

type MyStack struct {
	stack [100]int
	size  int
}

func Constructor_() MyStack {
	return MyStack{
		stack: [100]int{},
		size:  0,
	}
}

func (this *MyStack) Push(x int) {
	this.stack[this.size] = x
	this.size++

}

func (this *MyStack) Pop() int {
	if this.size == 0 {
		return -1
	}

	this.size--
	return this.stack[this.size]
}

func (this *MyStack) Top() int {
	return this.stack[this.size-1]
}

func (this *MyStack) Empty() bool {
	return this.size == 0
}
