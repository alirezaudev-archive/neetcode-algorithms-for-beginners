package main

// https://leetcode.com/problems/binary-tree-right-side-view/

import (
	"../../libs/Go/tree"
	"fmt"
)

func main() {
	root := tree.FromArray([]int{1, 2, 3, -1, 5, -1, 4})
	root.Print()
	fmt.Printf("%v\n", rightSideView(root))
}

func rightSideView(root *tree.Node) []int {
	var result []int
	if root == nil {
		return result
	}
	queue := []*tree.Node{root}

	for len(queue) > 0 {
		var rightView int
		queueLen := len(queue)
		for i := 0; i < queueLen; i++ {
			curr := queue[i]
			rightView = curr.Val
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		queue = queue[queueLen:]
		result = append(result, rightView)
	}

	return result
}
