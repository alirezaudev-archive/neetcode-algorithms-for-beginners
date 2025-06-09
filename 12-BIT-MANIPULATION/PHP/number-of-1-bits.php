<?php

// https://leetcode.com/problems/number-of-1-bits/

class Solution
{
    function hammingWeight(int $n): int
    {
        $c = 0;
        while ($n > -0) {
            if ($n & 1) {
                $c++;
            }

            $n >>= 1;
        }

        return $c;
    }
}