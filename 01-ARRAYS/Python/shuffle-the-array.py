from typing import *


# https://leetcode.com/problems/shuffle-the-array/
def shuffle(nums: List[int], n: int) -> List[int]:
    for i in range(n):
        nums.insert(i * 2 + 1, nums.pop(i + n))
    return nums
