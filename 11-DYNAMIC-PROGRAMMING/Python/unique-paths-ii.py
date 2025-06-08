from typing import List

# https://leetcode.com/problems/unique-paths-ii/

class Solution:
    def uniquePathsWithObstacles(self, obstacleGrid: List[List[int]]) -> int:
        rows, cols = len(grid), len(grid[0])
        memo = {(rows-1, cols-1): 1}

        def dfs(r: int, c: int) -> int:
            if r == rows or c == cols or grid[r][c]:
                return 0
            point = (r, c)
            if point in memo:
                return memo[point]

            memo[point] = dfs(r + 1, c) + dfs(r, c + 1)
            return memo[point]

        return dfs(0, 0)