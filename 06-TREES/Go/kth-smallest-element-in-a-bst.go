package main

// https://leetcode.com/problems/kth-smallest-element-in-a-bst/

import (
	"../../libs/Go/tree"
	"fmt"
)

func main() {
	root := tree.FromArray([]int{3, 1, 4, -1, 2})
	root.Print()
	fmt.Printf("k: %d\n", kthSmallest(root, 1))
}

func kthSmallest(root *tree.Node, k int) int {
	return trav(root, &k)
}

func trav(root *tree.Node, k *int) int {
	if root == nil {
		return -1
	}

	res := trav(root.Left, k)
	if res != -1 {
		return res
	}

	*k--
	if *k == 0 {
		return root.Val
	}

	return trav(root.Right, k)
}
