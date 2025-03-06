package main

// https://leetcode.com/problems/kth-largest-element-in-an-array/

import (
	"container/heap"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &MinHeap{}
	heap.Init(h)

	for i := 0; i < k; i++ {
		heap.Push(h, nums[i])
	}

	for i := k; i < len(nums); i++ {
		if nums[i] > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, nums[i])
		}
	}

	return (*h)[0]
}

func findKthLargest2(nums []int, k int) int {
	stack := nums[:k]
	for i, v := range nums[k:] {
		i += k
		min_, minI := findMin(stack)
		if v > min_ {
			stack[minI] = v
		}
	}

	kthLargest, _ := findMin(stack)
	return kthLargest
}

func findMin(stack []int) (min_, minI int) {
	min_ = stack[0]
	minI = 0
	for j, m := range stack[1:] {
		if m < min_ {
			min_ = m
			minI = j + 1
		}
	}

	return
}

func findUsingQuickSort(nums []int, start, end, k int) int {
	if start == end {
		return nums[start]
	}

	pivot := nums[end]
	pointer := start

	for i := start; i < end; i++ { // Fix: iterate only up to end-1
		if nums[i] < pivot {
			nums[pointer], nums[i] = nums[i], nums[pointer]
			pointer++
		}
	}

	nums[pointer], nums[end] = nums[end], nums[pointer]

	if pointer == k {
		return nums[pointer]
	} else if pointer > k {
		return findUsingQuickSort(nums, start, pointer-1, k)
	} else {
		return findUsingQuickSort(nums, pointer+1, end, k)
	}
}
