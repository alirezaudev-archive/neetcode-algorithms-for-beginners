package main

// https://leetcode.com/problems/last-stone-weight/

func lastStoneWeight(stones []int) int {
	heap := append([]int{}, stones...)

	for i := len(heap)/2 - 1; i >= 0; i-- {
		heap = heapify(heap, i)
	}

	for len(heap) > 1 {
		y := heap[0]
		heap = popHeap(heap)

		if len(heap) == 0 {
			return y
		}

		x := heap[0]
		heap = popHeap(heap)

		diff := y - x
		if diff > 0 {
			heap = pushHeap(heap, diff)
		}
	}

	if len(heap) == 0 {
		return 0
	}

	return heap[0]
}

func pushHeap(heap []int, stone int) []int {
	heap = append(heap, stone)
	i := len(heap) - 1
	for i > 0 && heap[(i-1)/2] < heap[i] {
		parent := (i - 1) / 2
		heap[parent], heap[i] = heap[i], heap[parent]
		i = parent
	}

	return heap
}

func popHeap(heap []int) []int {
	length := len(heap)
	if length < 1 {
		return []int{}
	}

	last := length - 1
	heap[0] = heap[last]
	return heapify(heap[:last], 0)
}

func maxChild(heap []int, i int) int {
	length := len(heap)
	left := i*2 + 1
	right := i*2 + 2

	if right < length && heap[right] > heap[left] {
		return right
	}

	if left < length {
		return left
	}

	return i
}

func heapify(heap []int, i int) []int {
	m := maxChild(heap, i)
	for m < len(heap) && heap[i] < heap[m] {
		heap[i], heap[m] = heap[m], heap[i]
		i = m
		m = maxChild(heap, i)
	}
	return heap
}
