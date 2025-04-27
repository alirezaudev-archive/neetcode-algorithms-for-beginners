<?php

// https://leetcode.com/problems/number-of-islands/

class Solution
{
    private array $grid = [];
    private int $rows = 0;
    private int $cols = 0;

    /**
     * @param string[][] $grid
     * @return int
     */
    function numIslands(array $grid): int
    {
        if (count($grid) == 0) {
            return 0;
        }

        $this->grid = $grid;
        $this->rows = count($grid);
        $this->cols = count($grid[0]);

        $count = 0;
        for ($r = 0; $r < $this->rows; $r++) {
            for ($c = 0; $c < $this->cols; $c++) {
                if ($this->grid[$r][$c] == '1') {
                    $count++;
                    $this->dfs($r, $c);
                }
            }
        }
        return $count;
    }

    private function dfs(int $r, int $c): void
    {
        if ($this->isOutOfBound($r, $c) || $this->grid[$r][$c] == '0') {
            return;
        }

        $this->grid[$r][$c] = '0';
        $this->dfs($r-1, $c);
        $this->dfs($r+1, $c);
        $this->dfs($r, $c-1);
        $this->dfs($r, $c+1);
    }

    private function isOutOfBound(int $r, int $c): bool
    {
        return min($r, $c) < 0 || $r >= $this->rows || $c >= $this->cols;
    }

}