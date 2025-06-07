package main

// https://leetcode.com/problems/house-robber/

func rob(nums []int) int {
	rob1, rob2 := 0, 0
	for _, n := range nums {
		tmp := max(n+rob1, rob2)
		rob1 = rob2
		rob2 = tmp
	}

	return rob2
}
