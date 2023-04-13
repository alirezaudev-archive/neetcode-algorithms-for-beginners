<?php

# https://leetcode.com/problems/sort-an-array/
class Solution
{
    /**
     * Insertion Sorting Algorithm
     */
    public function insertionSortArray(array $nums): array
    {
        for ($i = 1; $i < count($nums); $i++) {
            $j = $i - 1;
            while ($j >= 0 && $nums[$j + 1] < $nums[$j]) {
                $tmp = $nums[$j];
                $nums[$j] = $nums[$j + 1];
                $nums[$j + 1] = $tmp;
                $j--;
            }

        }
        return $nums;
    }
}