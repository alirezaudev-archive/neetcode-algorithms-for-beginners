package main

// https://leetcode.com/problems/sort-colors/
func sortColors(nums []int) {
	buckets := [3]int{}

	for _, i := range nums {
		buckets[i]++
	}

	i := 0
	for num, bucket := range buckets {
		for count := 0; count < bucket; count++ {
			nums[i] = num
			i++
		}
	}
}
