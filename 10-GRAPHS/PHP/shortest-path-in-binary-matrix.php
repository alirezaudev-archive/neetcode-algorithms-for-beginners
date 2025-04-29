<?php

// https://leetcode.com/problems/shortest-path-in-binary-matrix/
class Solution
{
    /**
     * @param Integer[][] $grid
     * @return Integer
     */
    function shortestPathBinaryMatrix(array $grid): int
    {
        $n = count($grid);
        if ($grid[0][0] !== 0 || $grid[$n - 1][$n - 1] !== 0) {
            return -1;
        }

        $directions = [[-1, -1], [-1, 0], [-1, 1], [0, -1], [0, 1], [1, -1], [1, 0], [1, 1]];

        $queue = [[0, 0]];
        $head = 0;
        $grid[0][0] = 1;

        while ($head < count($queue)) {
            [$row, $col] = $queue[$head++];
            $distance = $grid[$row][$col];

            if ($row === $n - 1 && $col === $n - 1) {
                return $distance;
            }

            foreach ($directions as [$dr, $dc]) {
                $rn = $row + $dr;
                $cn = $col + $dc;

                if (
                    $rn >= 0 && $rn < $n &&
                    $cn >= 0 && $cn < $n &&
                    $grid[$rn][$cn] === 0
                ) {
                    $grid[$rn][$cn] = $distance + 1;
                    $queue[] = [$rn, $cn];
                }
            }
        }

        return -1;
    }
}
