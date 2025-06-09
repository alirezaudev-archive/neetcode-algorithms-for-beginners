<?php

// https://leetcode.com/problems/counting-bits/

class Solution
{
    function countBits(int $n): array
    {
        $bits = array_fill(0, $n+1, 0);
        for ($i = 1; $i <= $n; $i++) {
            $bits[$i] = $bits[$i >> 1] + ($i & 1);
        }

        return $bits;
    }
}