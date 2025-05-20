<?php

// https://leetcode.com/problems/rotting-oranges/

class Solution
{
    /**
     * @param Integer[][] $grid
     * @return Integer
     */
    function orangesRotting(array $grid): int
    {
        $rows = count($grid);
        if ($rows === 0) {
            return 0;
        }

        $cols = count($grid[0]);
        $directions = [[-1, 0], [1, 0], [0, 1], [0, -1]];
        $queue = [];
        $head = 0;
        $fresh = 0;
        $result = -1;

        for ($r = 0; $r < $rows; $r++) {
            for ($c = 0; $c < $cols; $c++) {
                if ($grid[$r][$c] === 1) {
                    $fresh++;
                } else if ($grid[$r][$c] === 2) {
                    $queue[] = [$r, $c];
                }
            }
        }

        if ($fresh === 0) {
            return 0;
        }

        while ($head < count($queue)) {
            $size = count($queue) - $head;
            $result++;

            for ($i = 0; $i < $size; $i++) {
                [$r, $c] = $queue[$head++];
                foreach ($directions as [$dr, $dc]) {
                    $rn = $r + $dr;
                    $cn = $c + $dc;

                    if (
                        $rn >= 0 && $rn < $rows &&
                        $cn >= 0 && $cn < $cols &&
                        $grid[$rn][$cn] === 1
                    ) {
                        $grid[$rn][$cn] = 2;
                        $fresh--;
                        $queue[] = [$rn, $cn];
                    }
                }
            }
        }

        return $fresh > 0 ? -1 : $result;
    }
}
