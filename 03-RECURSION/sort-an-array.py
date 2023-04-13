from typing import *


# https://leetcode.com/problems/sort-an-array/
class Solution:
    def insertionSortArray(self, nums: List[int]) -> List[int]:
        """
        Insertion Sorting Algorithm
        """
        for i in range(1, len(nums)):
            j = i - 1
            while j >= 0 and nums[j + 1] < nums[j]:
                nums[j], nums[j + 1] = nums[j + 1], nums[j]
                j -= 1

        return nums
