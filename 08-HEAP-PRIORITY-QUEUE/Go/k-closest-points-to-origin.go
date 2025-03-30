package main

// https://leetcode.com/problems/k-closest-points-to-origin/

type heap struct {
	k      int
	points [][]int
}

func kClosest(points [][]int, k int) [][]int {
	h := heap{k: k, points: [][]int{}}
	for _, point := range points {
		h.add(point)
	}

	for i := range h.points {
		h.points[i] = h.points[i][:2]
	}

	return h.points
}

func (this *heap) add(point []int) {
	point = append(point, calcDistance(point))
	length := len(this.points)
	if length < this.k {
		this.points = append(this.points, point)
		i := length
		for this.points[(i-1)/2][2] < this.points[i][2] {
			parent := (i - 1) / 2
			this.points[i], this.points[parent] = this.points[parent], this.points[i]
			i = parent
		}
	} else if point[2] < this.points[0][2] {
		this.points[0] = point
		i := 0
		for {
			largest := i
			left := 2*i + 1
			right := 2*i + 2
			if left < this.k && this.points[largest][2] < this.points[left][2] {
				largest = left
			}

			if right < this.k && this.points[largest][2] < this.points[right][2] {
				largest = right
			}

			if largest == i {
				break
			}

			this.points[i], this.points[largest] = this.points[largest], this.points[i]
			i = largest
		}
	}
}

func calcDistance(p1 []int) int {
	return p1[0]*p1[0] + p1[1]*p1[1]
}
