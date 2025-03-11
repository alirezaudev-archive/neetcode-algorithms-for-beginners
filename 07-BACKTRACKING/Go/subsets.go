package main

import "fmt"

// https://leetcode.com/problems/subsets/

func main() {
	fmt.Printf("%v\n", subsets([]int{1, 2, 3}))
	fmt.Printf("%v\n", subsets2([]int{1, 2, 3}))
}

func subsets2(nums []int) [][]int {
	var res [][]int
	var temp []int

	var backtrack func(index int)
	backtrack = func(index int) {
		if index >= len(nums) {
			subsetCopy := make([]int, len(temp))
			copy(subsetCopy, temp)
			res = append(res, subsetCopy)
			return
		}

		temp = append(temp, nums[index])
		backtrack(index + 1)
		temp = temp[:len(temp)-1]
		backtrack(index + 1)
	}

	backtrack(0)
	return res
}

func subsets(nums []int) [][]int {
	subs := [][]int{{}}

	for _, num := range nums {
		var tmp [][]int
		for _, sub := range subs {
			newSub := append([]int{}, sub...)
			tmp = append(tmp, append(newSub, num))
		}
		subs = append(subs, tmp...)
	}

	return subs
}
