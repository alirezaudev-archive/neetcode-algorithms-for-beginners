package main

// https://leetcode.com/problems/insert-into-a-binary-search-tree/

import "../../libs/Go/tree"

func insertIntoBST(root *tree.Node, val int) *tree.Node {
	if root == nil {
		return &tree.Node{Val: val}
	}

	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}

	return root
}
