<?php

# https://leetcode.com/problems/binary-search/
class Solution
{
    public function search(array $nums, int $target): int
    {
        $low = 0;
        $high = count($nums) - 1;

        while ($low <= $high) {
            $mid = floor(($low + $high) / 2);

            if ($target > $nums[$mid]) {
                $low = $mid + 1;
            } elseif ($target < $nums[$mid]) {
                $high = $mid - 1;
            } else {
                return $mid;
            }
        }
        return -1;
    }
}