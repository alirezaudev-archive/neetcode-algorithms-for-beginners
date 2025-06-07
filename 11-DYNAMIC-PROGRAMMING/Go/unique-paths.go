package main

// https://leetcode.com/problems/unique-paths/

func uniquePaths(m int, n int) int {
	memo := make(map[[2]int]int, n*m*2)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			point := [2]int{i, j}
			if j == n-1 || i == m-1 {
				memo[point] = 1
				continue
			}

			memo[point] = memo[[2]int{i + 1, j}] + memo[[2]int{i, j + 1}]
		}
	}

	return memo[[2]int{0, 0}]
}
