<?php

// https://leetcode.com/problems/house-robber/

class Solution
{

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function rob(array $nums): int
    {
        $rob1 = 0;
        $rob2 = 0;
        foreach ($nums as $num) {
            $tmp = max($rob1+$num, $rob2);
            $rob1 = $rob2;
            $rob2 = $tmp;
        }

        return $rob2;
    }
}