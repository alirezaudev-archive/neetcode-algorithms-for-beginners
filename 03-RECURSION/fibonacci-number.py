from typing import *


# https://leetcode.com/problems/fibonacci-number/
class Solution:
    def fib(self, n: int) -> int:
        if n == 0 or n == 1:
            return n

        return self.fib(n - 1) + self.fib(n - 2)

    """
    This solution is unbelievably fast! 🚀😄
    """
    def fibFaster(self, n: int) -> int:
        fibs = (0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811, 514229, 832040)

        return fibs[n]