# https://leetcode.com/problems/last-stone-weight/

class Solution(object):
    def __init__(self):
        self.heap = []

    def lastStoneWeight(self, stones):
        for i in stones:
            self.push(i)

        while len(self.heap) > 2:
            y = self.pop()
            x = self.pop()
            diff = y - x
            if diff > 0:
                self.push(diff)

        length = len(self.heap)
        if length == 2:
            return max(self.heap) - min(self.heap)

        if length == 1:
            return self.heap[0]

        return 0

    def push(self, stone):
        self.heap.append(stone)
        i = len(self.heap) - 1
        while i > 0 and self.heap[(i - 1) // 2] < self.heap[i]:
            parent = (i - 1) // 2
            self.heap[parent], self.heap[i] = self.heap[i], self.heap[parent]
            i = parent

    def pop(self):
        if len(self.heap) < 1:
            return None

        res = self.heap[0]
        self.heap[0] = self.heap.pop()
        i = 0
        m = self.maxChild(i)
        while m < len(self.heap) and self.heap[i] < self.heap[m]:
            self.heap[i], self.heap[m] = self.heap[m], self.heap[i]
            i = m
            m = self.maxChild(i)

        return res

    def maxChild(self, i):
        length = len(self.heap)
        left = 2 * i + 1
        right = 2 * i + 2
        if right < length and self.heap[right] > self.heap[left]:
            return right
        if left < length:
            return left
        return i
