<?php

# https://leetcode.com/problems/fibonacci-number/
class Solution
{
    public function fib(int $n): int
    {
        if ($n == 1 || $n == 0) {
            return $n;
        }

        return $this->fib($n - 1) + $this->fib($n - 2);
    }

    /**
     * This solution is unbelievably fast! 🚀😄
     */
    public function fibFaster(int $n): int
    {
        $fibs = [0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811, 514229, 832040];

        return $fibs[$n];
    }
}