from typing import *


# https://leetcode.com/problems/sort-colors/
class Solution:
    def sortColors(self, nums: List[int]) -> None:
        """
        Bucket Sorting Algorithm
        """
        counts = [0, 0, 0]

        for n in nums:
            counts[n] += 1

        i = 0
        for n in range(len(counts)):
            for _ in range(counts[n]):
                nums[i] = n
                i += 1

