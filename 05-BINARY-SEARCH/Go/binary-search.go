package main

// https://leetcode.com/problems/binary-search/
func search(nums []int, target int) int {
	return binSearch(nums, target, 0, len(nums)-1)
}

func binSearch(nums []int, target, start, end int) int {
	if start > end {
		return -1
	}

	mid := (start + end) / 2
	if nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		return binSearch(nums, target, mid+1, end)
	} else {
		return binSearch(nums, target, start, mid-1)
	}
}
