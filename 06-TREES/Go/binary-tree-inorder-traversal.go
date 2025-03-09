package main

import "../../libs/Go/tree"

// https://leetcode.com/problems/binary-tree-inorder-traversal/
func inorderTraversal(root *tree.Node) []int {
	var result []int
	return traverse(root, result)
}

func traverse(root *tree.Node, result []int) []int {
	if root != nil {
		result = traverse(root.Left, result)
		result = append(result, root.Val)
		result = traverse(root.Right, result)
	}
	return result
}
