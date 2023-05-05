<?php


# https://leetcode.com/problems/sort-colors/
class Solution
{
    /**
     * Bucket Sorting Algorithm
     */
    public function sortColors(array &$nums): void
    {
        $counts = [0, 0, 0];

        foreach ($nums as $n) {
            $counts[$n]++;
        }

        $i = 0;
        for ($n = 0; $n < count($counts); $n++) {
            for ($j = 0; $j < $counts[$n]; $j++) {
                $nums[$i] = $n;
                $i++;
            }
        }
    }
}