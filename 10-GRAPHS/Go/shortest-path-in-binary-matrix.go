package main

// https://leetcode.com/problems/shortest-path-in-binary-matrix/

const N = 10000

func shortestPathBinaryMatrix(grid [][]int) int {
	if len(grid) == 0 || grid[0][0] == 1 {
		return -1
	}

	queue := deque{}
	visited := make(map[[2]int]bool, N)
	queue.push([2]int{0, 0})
	visited[[2]int{0, 0}] = true

	rows, cols := len(grid), len(grid[0])
	if grid[rows-1][cols-1] == 1 {
		return -1
	}

	result := 1
	for !queue.empty() {
		size := queue.size
		for i := 0; i < size; i++ {
			point := queue.pop()
			r, c := point[0], point[1]
			if r == rows-1 && c == cols-1 {
				return result
			}

			directions := [8][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}, {-1, -1}, {1, 1}, {-1, 1}, {1, -1}}
			for _, dir := range directions {
				rn, cn := r+dir[0], c+dir[1]
				newPoint := [2]int{rn, cn}
				if min(rn, cn) < 0 || rn >= rows || cn >= cols || visited[newPoint] || grid[rn][cn] == 1 {
					continue
				}

				visited[newPoint] = true
				queue.push(newPoint)
			}
		}
		result++
	}

	return -1
}

type deque struct {
	queue [N][2]int
	head  int
	tail  int
	size  int
}

func (this *deque) push(x [2]int) {
	if this.size == N {
		return
	}
	this.queue[this.tail] = x
	this.tail++
	if this.tail == N {
		this.tail = 0
	}
	this.size++
}

func (this *deque) pop() [2]int {
	if this.size == 0 {
		return [2]int{-1, -1}
	}
	val := this.queue[this.head]
	this.head++
	if this.head == N {
		this.head = 0
	}
	this.size--
	return val
}

func (this *deque) top() [2]int {
	if this.size == 0 {
		return [2]int{-1, -1}
	}
	return this.queue[this.head]
}

func (this *deque) empty() bool {
	return this.size == 0
}
