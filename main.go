package main

import (
	"fmt"
)

func main() {
	obj := Constructor(3, []int{7, 7, 7, 7, 8, 3})
	fmt.Printf("nums: %v\n", obj.nums)
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
