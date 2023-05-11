from typing import *


# https://leetcode.com/problems/search-a-2d-matrix/
class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        for row in matrix:
            if not (target >= row[0] or target <= row[-1]):
                continue

            low = 0
            high = len(row) - 1

            while low <= high:
                mid = (low + high) // 2

                if target > row[mid]:
                    low = mid + 1
                elif target < row[mid]:
                    high = mid - 1
                else:
                    return True

        return False
