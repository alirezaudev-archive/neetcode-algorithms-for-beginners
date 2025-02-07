package main

func shuffle(nums []int, n int) []int {
	shuffled := make([]int, n*2)
	for i := 0; i < n*2; i += 2 {
		shuffled[i] = nums[i/2]
		shuffled[i+1] = nums[(i/2)+n]
	}

	return shuffled
}
