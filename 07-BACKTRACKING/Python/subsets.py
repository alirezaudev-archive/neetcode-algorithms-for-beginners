from typing import *
from monitoring import execute

# https://leetcode.com/problems/subsets/
class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        subs = [[]]

        for num in nums:
            temp = [sub + [num] for sub in subs]
            subs.append(temp)

        return subs

    def backtrackSubs(self, nums: List[int]) -> List[List[int]]:
        """
        using backtrack algorithm
        :param nums:
        :return:
        """
        def backtrack(start, sub):
            subs.append(sub[:])
            for i in range(start, len(nums)):
                sub.append(nums[i])
                backtrack(i + 1, sub)
                sub.pop()

        subs = []
        backtrack(0, [])
        return subs


lst = list(range(20))
solution = Solution()


print('LOOP')
execute(solution.subsets, lst)

print()

print('BACKTRACKING')
execute(solution.backtrackSubs, lst)