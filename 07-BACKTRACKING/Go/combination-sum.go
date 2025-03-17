package main

// https://leetcode.com/problems/combination-sum/
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int

	var backtrack func(index int, curr []int, total int)
	backtrack = func(index int, curr []int, total int) {
		if total == target {
			currCopy := make([]int, len(curr))
			copy(currCopy, curr)
			res = append(res, currCopy)
			return
		}

		if total > target || index >= len(candidates) {
			return
		}

		curr = append(curr, candidates[index])
		backtrack(index, curr, total+candidates[index])
		curr = curr[:len(curr)-1]
		backtrack(index+1, curr, total)
	}

	backtrack(0, []int{}, 0)
	return res
}
