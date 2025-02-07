from typing import *


# https://leetcode.com/problems/remove-duplicates-from-sorted-array/
def removeDuplicates(nums: List[int]) -> int:
    uniquesLength = 1
    for i in range(1, len(nums)):
        if nums[i - 1] != nums[i]:
            nums[uniquesLength] = nums[i]
            uniquesLength += 1

    return uniquesLength
