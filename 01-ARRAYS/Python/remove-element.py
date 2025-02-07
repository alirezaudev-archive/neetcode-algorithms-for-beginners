from typing import *


# https://leetcode.com/problems/remove-element/
def removeElement(nums: List[int], val: int) -> int:
    k = 0
    for num in nums:
        if num != val:
            nums[k] = num
            k += 1
    return k


