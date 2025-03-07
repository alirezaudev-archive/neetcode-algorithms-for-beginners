package main

import "../../libs/Go/tree"

// https://leetcode.com/problems/search-in-a-binary-search-tree/
func searchBST(root *tree.Node, val int) *tree.Node {
	if root == nil || root.Val == val {
		return root
	}

	if root.Val < val {
		return searchBST(root.Right, val)
	}

	return searchBST(root.Left, val)
}
