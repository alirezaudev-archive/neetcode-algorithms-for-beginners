<?php

// https://leetcode.com/problems/two-sum/

class Solution
{
    function twoSum(array $nums, int $target): array
    {
        $l = count($nums);
        for ($i = 0; $i < $l; $i++) {
            for ($j = $i + 1; $j < $l; $j++) {
                if ($nums[$i] + $nums[$j] == $target) {
                    return [$i, $j];
                }
            }
        }

        return [];
    }
}