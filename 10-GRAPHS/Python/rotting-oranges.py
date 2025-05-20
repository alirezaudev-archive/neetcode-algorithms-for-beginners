from collections import deque

# https://leetcode.com/problems/rotting-oranges/

class Solution:
    def orangesRotting(self, grid: List[List[int]]) -> int:
        if not grid:
            return 0

        rows, cols = len(grid), len(grid[0])
        fresh = 0
        queue = deque()

        for r in range(rows):
            for c in range(cols):
                if grid[r][c] == 1:
                    fresh += 1
                elif grid[r][c] == 2:
                    queue.append((r, c))

        if fresh == 0:
            return 0

        directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]
        result = -1

        while queue:
            level_size = len(queue)
            result += 1
            for _ in range(level_size):
                r, c = queue.popleft()
                for dr, dc in directions:
                    rn, cn = r + dr, c + dc
                    if 0 <= rn < rows and 0 <= cn < cols and grid[rn][cn] == 1:
                        grid[rn][cn] = 2
                        fresh -= 1
                        queue.append((rn, cn))

        return -1 if fresh > 0 else result
