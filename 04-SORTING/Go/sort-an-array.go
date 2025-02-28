package main

// https://leetcode.com/problems/sort-an-array/
func sortArray(nums []int) []int {
	return mergeSortArray(nums, 0, len(nums)-1)
}

func mergeSortArray(nums []int, start, end int) []int {
	if start >= end {
		return nums[start : end+1]
	}

	middle := (start + end) / 2

	return mergeTwoArray(
		mergeSortArray(nums, start, middle),
		mergeSortArray(nums, middle+1, end),
	)
}

func mergeTwoArray(left, right []int) []int {
	leftLen := len(left)
	rightLen := len(right)

	merged := make([]int, leftLen+rightLen)

	var i, l, r int
	for l < leftLen && r < rightLen {
		if left[l] <= right[r] {
			merged[i] = left[l]
			l++
		} else {
			merged[i] = right[r]
			r++
		}
		i++
	}

	for l < leftLen {
		merged[i] = left[l]
		l++
		i++
	}

	for r < rightLen {
		merged[i] = right[r]
		r++
		i++
	}

	return merged
}

// sortArray insertion sort algorithm O(n^2) | Time Limit Exceeded
func insertionSortArray(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j > 0 && nums[j] < nums[j-1]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}

	return nums
}
