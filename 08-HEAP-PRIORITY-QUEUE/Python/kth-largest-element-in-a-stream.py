from typing import List

# https://leetcode.com/problems/kth-largest-element-in-a-stream/

class KthLargest:
    def __init__(self, k: int, nums: List[int]):
        self.k = k
        self.heap = []

        for i in nums:
            self.add(i)

    def add(self, val: int) -> int:
        length = len(self.heap)
        if length < self.k:
            self.heap.append(val)
            i = length
            while i > 0:
                parent = (i - 1) // 2
                if self.heap[parent] <= self.heap[i]:
                    break
                self.heap[parent], self.heap[i] = self.heap[i], self.heap[parent]
                i = parent
        elif val > self.heap[0]:
            self.heap[0] = val
            i = 0
            n = self.k

            while True:
                smallest = i
                left = 2 * i + 1
                right = 2 * i + 2
                if left < n and self.heap[left] < self.heap[smallest]:
                    smallest = left

                if right < n and self.heap[right] < self.heap[smallest]:
                    smallest = right

                if smallest == i:
                    break

                self.heap[i], self.heap[smallest] = self.heap[smallest], self.heap[i]
                i = smallest

        return self.heap[0]
