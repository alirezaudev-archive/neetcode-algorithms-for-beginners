from typing import List


# https://leetcode.com/problems/number-of-islands/
class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        if len(grid) == 0:
            return 0

        rows, cols = len(grid), len(grid[0])

        def dfs(r: int, c: int):
            if min(r, c) < 0 or r >= rows or c >= cols or grid[r][c] == '0':
                return

            grid[r][c] = '0'
            dfs(r - 1, c)
            dfs(r + 1, c)
            dfs(r, c - 1)
            dfs(r, c + 1)

        result = 0
        for r in range(rows):
            for c in range(cols):
                if grid[r][c] == '1':
                    result += 1
                    dfs(r, c)

        return result
