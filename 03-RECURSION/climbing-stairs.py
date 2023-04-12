from typing import *


# https://leetcode.com/problems/climbing-stairs/
class Solution:
    def climbStairs(self, n: int) -> int:
        if n <= 1:
            return 1

        return self.climbStairs(n - 1) + self.climbStairs(n - 2)

    def climbStairsFaster(self, n: int) -> int:
        memo = {}

        return self.helper(n, memo)

    def helper(self, n: int, memo: dict) -> int:
        if n <= 1:
            return 1

        if n not in memo:
            memo[n] = self.helper(n - 1, memo) + self.helper(n - 2, memo)

        return memo[n]
