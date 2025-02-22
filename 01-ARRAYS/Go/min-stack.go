package main

type MinStack struct {
	stack          [30000]int
	length         int
	minimums       [30000]int
	minimumsLength int
}

func Constructor() MinStack {
	return MinStack{
		length:         0,
		minimumsLength: 0,
	}
}

func (this *MinStack) Push(val int) {
	this.stack[this.length] = val
	if this.length == 0 || this.minimums[this.minimumsLength-1] >= val {
		this.minimums[this.minimumsLength] = val
		this.minimumsLength++
	}

	this.length++
}

func (this *MinStack) Pop() {
	if this.length == 0 {
		return
	}

	if this.minimumsLength > 0 && this.Top() == this.minimums[this.minimumsLength-1] {
		this.minimumsLength--
	}

	this.length--
}

func (this *MinStack) Top() int {
	return this.stack[this.length-1]
}

func (this *MinStack) GetMin() int {
	return this.minimums[this.minimumsLength-1]
}
