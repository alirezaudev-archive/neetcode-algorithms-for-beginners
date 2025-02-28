package main

// https://leetcode.com/problems/climbing-stairs/
func climbStairs(n int) int {
	memo := make(map[int]int, 45)
	return calc(n, memo)
}

func calc(n int, memo map[int]int) int {
	if n < 2 {
		return 1
	}

	if _, ok := memo[n]; !ok {
		memo[n] = calc(n-1, memo) + calc(n-2, memo)
	}

	return memo[n]
}
