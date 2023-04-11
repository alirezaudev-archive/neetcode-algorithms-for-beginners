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
}