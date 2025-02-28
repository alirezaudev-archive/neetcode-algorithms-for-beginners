package main

// https://leetcode.com/problems/fibonacci-number/
func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
