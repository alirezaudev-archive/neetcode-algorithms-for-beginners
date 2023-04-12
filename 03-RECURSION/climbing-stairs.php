<?php


# https://leetcode.com/problems/climbing-stairs/
class Solution
{
    public function climbStairs(int $n): int
    {
        if ($n <= 1) {
            return 1;
        }

        return $this->climbStairs($n - 1) + $this->climbStairs($n - 2);
    }

    public function climbStairsFaster(int $n): int
    {
        $memo = [];

        return $this->helper($n, $memo);
    }

    public function helper(int $n, array $memo): int
    {
        if ($n <= 1) {
            return 1;
        }

        if (! in_array($n, $memo)) {
            $memo[$n] = $this->helper($n - 1, $memo) + $this->helper($n - 2, $memo);
        }

        return $memo[$n];
    }
}
