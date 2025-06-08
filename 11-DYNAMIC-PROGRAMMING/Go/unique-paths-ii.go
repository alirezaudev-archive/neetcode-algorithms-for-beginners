package main

// https://leetcode.com/problems/unique-paths-ii/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	rows, cols := len(obstacleGrid), len(obstacleGrid[0])
	memo := make(map[[2]int]int, rows*cols*2)
	memo[[2]int{rows - 1, cols - 1}] = 1

	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r == rows || c == cols || obstacleGrid[r][c] == 1 {
			return 0
		}
		point := [2]int{r, c}
		if v, ok := memo[point]; ok {
			return v
		}

		memo[point] = dfs(r+1, c) + dfs(r, c+1)

		return memo[point]
	}

	return dfs(0, 0)
}

// NotDpUniquePathsWithObstacles my solution
func NotDpUniquePathsWithObstacles(obstacleGrid [][]int) int {
	rows, cols := len(obstacleGrid), len(obstacleGrid[0])
	memo := make(map[[2]int]int, rows*cols*2)

	if obstacleGrid[0][0] == 1 || obstacleGrid[rows-1][cols-1] == 1 {
		return 0
	}

	blockedCols := map[int]bool{}

	for i := rows - 1; i >= 0; i-- {
		thisRowHaveObstacle := false
		for j := cols - 1; j >= 0; j-- {
			point := [2]int{i, j}
			if obstacleGrid[i][j] == 1 {
				thisRowHaveObstacle = true
				blockedCols[j] = true
				continue
			}

			if j == cols-1 || i == rows-1 {
				if thisRowHaveObstacle || blockedCols[j] {
					continue
				}

				memo[point] = 1
				continue
			}

			memo[point] = memo[[2]int{i + 1, j}] + memo[[2]int{i, j + 1}]
		}
	}

	return memo[[2]int{0, 0}]
}
