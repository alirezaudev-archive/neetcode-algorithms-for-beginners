package main

import (
	"fmt"
)

// https://leetcode.com/problems/number-of-islands/

func main() {
	graph := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
	}

	fmt.Println(numIslands(graph))
}

var grid [][]byte
var cols, rows int
var visited map[[2]int]bool

func numIslands(grid_ [][]byte) int {
	if len(grid_) == 0 {
		return 0
	}
	grid = grid_
	rows, cols = len(grid), len(grid[0])
	visited = make(map[[2]int]bool, rows*cols)
	result := 0
	for r := range grid {
		for c := range grid[r] {
			point := [2]int{r, c}
			if grid[r][c] == '1' && !visited[point] {
				bfs(point)
				result++
			}
		}
	}
	return result
}

func bfs(point [2]int) {
	queue := deque{}
	queue.push(point)
	visited[point] = true
	directions := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for !queue.empty() {
		point = queue.pop()
		r, c := point[0], point[1]
		for _, direction := range directions {
			rn, cn := r+direction[0], c+direction[1]
			newPoint := [2]int{rn, cn}
			if isOutOFBound(rn, cn) || grid[rn][cn] == '0' || visited[newPoint] {
				continue
			}
			visited[newPoint] = true
			queue.push(newPoint)
		}
	}
}

func isOutOFBound(rn int, cn int) bool {
	return min(rn, cn) < 0 ||
		rn >= rows ||
		cn >= cols
}

type deque struct {
	stack [90000][2]int
	size  int
}

func (this *deque) push(x [2]int) {
	this.stack[this.size] = x
	this.size++
}

func (this *deque) pop() [2]int {
	if this.size == 0 {
		return [2]int{-1, -1}
	}

	this.size--
	return this.stack[this.size]
}

func (this *deque) top() [2]int {
	return this.stack[this.size-1]
}

func (this *deque) empty() bool {
	return this.size == 0
}
