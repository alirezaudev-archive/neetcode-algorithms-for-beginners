from typing import List

# https://leetcode.com/problems/max-area-of-island/
class Solution:
    def maxAreaOfIsland(self, grid: List[List[int]]) -> int:
        if len(grid) == 0:
            return 0

        rows, cols = len(grid), len(grid[0])

        def dfs(r: int, c: int) -> int:
            if min(r, c) < 0 or r >= rows or c >= cols or grid[r][c] == 0:
                return 0

            grid[r][c] = 0

            return 1 + dfs(r - 1, c) + dfs(r + 1, c) + dfs(r, c - 1) + dfs(r, c + 1)

        result = 0
        for r in range(rows):
            for c in range(cols):
                if grid[r][c] == 1:
                    result = max(result, dfs(r, c))

        return result
