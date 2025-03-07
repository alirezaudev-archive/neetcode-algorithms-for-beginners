package main

// https://leetcode.com/problems/concatenation-of-array/
func getConcatenation(nums []int) []int {
	numsLength := len(nums)
	result := make([]int, numsLength*2)
	for i, v := range nums {
		result[i] = v
		result[i+numsLength] = v
	}

	return result
}
