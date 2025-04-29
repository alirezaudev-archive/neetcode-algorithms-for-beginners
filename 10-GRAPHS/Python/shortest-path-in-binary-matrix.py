from collections import deque
from typing import *

# https://leetcode.com/problems/shortest-path-in-binary-matrix/

class Solution:
    def shortestPathBinaryMatrix(self, grid: List[List[int]]) -> int:
        if not grid or grid[0][0] == 1:
            return -1

        n = len(grid)
        if grid[n - 1][n - 1] == 1:
            return -1

        queue = deque()
        visited = set()

        queue.append((0, 0))
        visited.add((0, 0))
        result = 1

        directions = [(1, 0), (0, 1), (-1, 0), (0, -1),(-1, -1), (1, 1), (-1, 1), (1, -1)]

        while queue:
            for _ in range(len(queue)):
                r, c = queue.popleft()
                if r == n - 1 and c == n - 1:
                    return result

                for dr, dc in directions:
                    rn, cn = r + dr, c + dc
                    if (0 <= rn < n and 0 <= cn < n and
                            grid[rn][cn] == 0 and (rn, cn) not in visited):
                        visited.add((rn, cn))
                        queue.append((rn, cn))

            result += 1

        return -1
