package main

// https://leetcode.com/problems/delete-node-in-a-bst/

import "../../libs/Go/tree"

func main() {
	root := tree.FromArray([]int{5, 3, 6, 2, 4, -1, 7})
	root.Print()
	deleteNode(root, 4).Print()
	deleteNode(root, 6).Print()

}

func deleteNode(root *tree.Node, key int) *tree.Node {
	if root == nil {
		return root
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			minNode := root.Right
			for minNode.Left != nil {
				minNode = minNode.Left
			}
			root.Val = minNode.Val
			root.Right = deleteNode(root.Right, key)
		}
	}

	return root
}
