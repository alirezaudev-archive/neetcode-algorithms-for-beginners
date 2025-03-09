package main

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

import (
	"../../libs/Go/tree"
)

func main() {
	root := tree.FromArray([]int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5})
	root.Print()
	lowestCommonAncestor(root, root.Left.Left, root.Left.Right.Right).Print()
}

func lowestCommonAncestor(root, p, q *tree.Node) *tree.Node {
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return root
	}
}
