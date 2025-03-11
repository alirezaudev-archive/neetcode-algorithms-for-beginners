package main

//https://leetcode.com/problems/path-sum/

import "../../libs/Go/tree"

func main() {
	root := tree.FromArray([]int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, 1})
	root.Print()
	println(hasPathSum(root, 22))

}

func hasPathSum(root *tree.Node, targetSum int) bool {
	return hasPathSumHelper(root, &targetSum)
}

func hasPathSumHelper(root *tree.Node, targetSum *int) bool {
	if root == nil {
		return false
	}
	*targetSum -= root.Val

	if root.Left == nil && root.Right == nil && *targetSum == 0 {
		return true
	}

	if hasPathSumHelper(root.Left, targetSum) {
		return true
	}

	if hasPathSumHelper(root.Right, targetSum) {
		return true
	}

	*targetSum += root.Val

	return false
}
