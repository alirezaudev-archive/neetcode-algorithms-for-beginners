from heapq import heappush, heappop
from typing import List


# https://leetcode.com/problems/k-closest-points-to-origin/

class Solution:
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        max_heap = []

        for point in points:
            distance = -(point[0] ** 2 + point[1] ** 2)  # Use negative because Python's heapq is a min-heap
            heappush(max_heap, (distance, point))
            if len(max_heap) > k:
                heappop(max_heap)  # Remove the farthest point

        return [point for _, point in max_heap]
