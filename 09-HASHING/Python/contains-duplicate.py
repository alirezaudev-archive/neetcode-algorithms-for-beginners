from typing import List


# https://leetcode.com/problems/contains-duplicate/
class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        """
        I know I should've used a hashmap, but this was more fun.
        """
        return len(set(nums)) != len(nums)
