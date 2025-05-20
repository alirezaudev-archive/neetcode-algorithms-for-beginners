package main

// https://leetcode.com/problems/rotting-oranges/

func orangesRotting(grid_ [][]int) int {
	if len(grid_) == 0 {
		return 0
	}
	grid := grid_
	rows, cols := len(grid), len(grid[0])
	fresh := 0
	queue := deque{}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				fresh++
			} else if grid[r][c] == 2 {
				queue.push([2]int{r, c})
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	result := -1
	for !queue.empty() {
		levelSize := queue.size
		result++
		for i := 0; i < levelSize; i++ {
			point := queue.pop()
			r, c := point[0], point[1]

			for _, dir := range directions {
				rn, cn := r+dir[0], c+dir[1]
				if rn >= 0 && rn < rows && cn >= 0 && cn < cols && grid[rn][cn] == 1 {
					grid[rn][cn] = 2
					fresh--
					queue.push([2]int{rn, cn})
				}
			}
		}
	}

	if fresh > 0 {
		return -1
	}

	return result
}

const N = 100

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
