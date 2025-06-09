package main

// https://leetcode.com/problems/counting-bits/

func countBits(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
        bits[i] = bits[i>>1] + (i & 1)
    }
	return bits
}

func Solution1CountBits(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		n := i
		c := 0
		for n > 0 {
			c += n & 1
			n >>= 1
		}

		bits[i] = c
	}

	return bits
}
