package main

// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

import "../../libs/Go/tree"

func main() {
	buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}).Print()
}

func buildTree(preorder []int, inorder []int) *tree.Node {
	if len(preorder) == 0 {
		return nil
	}

	root := &tree.Node{Val: preorder[0]}
	rootIndex := findIndex(inorder, root.Val)
	root.Left = buildTree(preorder[1:rootIndex+1], inorder[0:rootIndex+1])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])

	return root
}

func findIndex(list []int, value int) int {
	for i, item := range list {
		if item == value {
			return i
		}
	}

	return -1
}
