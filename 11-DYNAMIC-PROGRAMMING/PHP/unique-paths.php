<?php

// https://leetcode.com/problems/unique-paths/

class Solution
{
    function uniquePaths(int $m, int $n): int
    {
        $row = array_fill(0, $n, 1);

        for ($i = 0; $i < $m - 1; $i++) {
            $newRow = array_fill(0, $n, 1);

            for ($j = $n - 2; $j >= 0; $j--) {
                $newRow[$j] = $newRow[$j + 1] + $row[$j];
            }

            $row = $newRow;
        }

        return $row[0];
    }
}
