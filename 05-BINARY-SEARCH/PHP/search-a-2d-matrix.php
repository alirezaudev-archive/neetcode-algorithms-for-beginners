<?php

# https://leetcode.com/problems/search-a-2d-matrix/
class Solution
{
    public function searchMatrix(array $matrix, int $target): bool
    {
        foreach ($matrix as $row) {
            if (!($target >= $row[0] || $target <= end($row))) {
                continue;
            }

            $low = 0;
            $high = count($row) - 1;

            while ($low <= $high) {
                $mid = floor(($low + $high) / 2);

                if ($target > $row[$mid]) {
                    $low = $mid + 1;
                } elseif ($target < $row[$mid]) {
                    $high = $mid - 1;
                } else {
                    return true;
                }
            }
        }

        return false;
    }
}