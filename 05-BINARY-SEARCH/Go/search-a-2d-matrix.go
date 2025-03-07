package main

// https://leetcode.com/problems/search-a-2d-matrix/
func searchMatrix(matrix [][]int, target int) bool {
	firstRow := 0
	lastRow := len(matrix) - 1
	var row int
	for firstRow <= lastRow {
		row = (firstRow + lastRow) / 2
		if target < matrix[row][0] {
			lastRow = row - 1
			continue
		}
		firstRow = row + 1

		if target <= matrix[row][len(matrix[row])-1] {
			left := 0
			right := len(matrix[row]) - 1
			for left <= right {
				mid := (left + right) / 2
				if target < matrix[row][mid] {
					right = mid - 1
				} else if target > matrix[row][mid] {
					left = mid + 1
				} else {
					return matrix[row][mid] == target
				}
			}
		}
	}

	return false
}
