package main

// https://leetcode.com/problems/binary-tree-level-order-traversal/

import (
	"../../libs/Go/tree"
	"fmt"
)

func main() {
	root := tree.FromArray([]int{3, 9, 20, -1, -1, 15, 7})
	root.Print()
	fmt.Printf("%v", levelOrder(root))
}

func levelOrder(root *tree.Node) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	queue := []*tree.Node{root}

	for len(queue) > 0 {
		var layer []int
		queueLen := len(queue)
		for i := 0; i < queueLen; i++ {
			curr := queue[i]
			layer = append(layer, curr.Val)
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		queue = queue[queueLen:]
		result = append(result, layer)
	}

	return result
}
