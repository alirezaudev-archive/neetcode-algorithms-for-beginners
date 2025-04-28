<?php

// https://leetcode.com/problems/max-area-of-island/
class Solution
{
    private array $grid = [];
    private int $rows = 0;
    private int $cols = 0;

    /**
     * @param int[][] $grid
     * @return int
     */
    function maxAreaOfIsland(array $grid): int
    {
        if (count($grid) == 0) {
            return 0;
        }

        $this->grid = $grid;
        $this->rows = count($grid);
        $this->cols = count($grid[0]);

        $result = 0;
        for ($r = 0; $r < $this->rows; $r++) {
            for ($c = 0; $c < $this->cols; $c++) {
                if ($this->grid[$r][$c] == 1) {
                    $result = max($result, $this->dfs($r, $c));
                }
            }
        }
        return $result;
    }

    private function dfs(int $r, int $c): int
    {
        if ($this->isOutOfBound($r, $c) || !$this->grid[$r][$c]) {
            return 0;
        }

        $this->grid[$r][$c] = 0;

        return 1
            + $this->dfs($r - 1, $c)
            + $this->dfs($r + 1, $c)
            + $this->dfs($r, $c - 1)
            + $this->dfs($r, $c + 1);
    }

    private function isOutOfBound(int $r, int $c): bool
    {
        return min($r, $c) < 0 || $r >= $this->rows || $c >= $this->cols;
    }

}