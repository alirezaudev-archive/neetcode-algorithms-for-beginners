package main

// https://leetcode.com/problems/kth-largest-element-in-a-stream/

type KthLargest struct {
	k    int
	nums []int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k, nums: []int{}}
	for _, num := range nums {
		kl.Add(num)
	}
	return kl
}

func (this *KthLargest) Add(val int) int {
	length := len(this.nums)
	if length < this.k {
		this.nums = append(this.nums, val)
		i := length
		for this.nums[(i-1)/2] > this.nums[i] {
			parent := (i - 1) / 2
			this.nums[i], this.nums[parent] = this.nums[parent], this.nums[i]
			i = parent
		}

	} else if val > this.nums[0] {
		this.nums[0] = val
		i := 0
		for {
			smallest := i
			left := 2*i + 1
			right := 2*i + 2
			if left < this.k && this.nums[smallest] > this.nums[left] {
				smallest = left
			}

			if right < this.k && this.nums[smallest] > this.nums[right] {
				smallest = right
			}

			if smallest == i {
				break
			}

			this.nums[i], this.nums[smallest] = this.nums[smallest], this.nums[i]
			i = smallest
		}
	}

	return this.nums[0]
}
