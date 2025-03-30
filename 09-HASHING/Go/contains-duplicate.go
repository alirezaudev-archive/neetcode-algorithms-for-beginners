package main

// https://leetcode.com/problems/contains-duplicate/

func containsDuplicate(nums []int) bool {
	h := map[int]bool{}
	for _, num := range nums {
		if h[num] {
			return true
		}
		h[num] = true
	}
	return false
}
