<?php

// https://leetcode.com/problems/unique-paths-ii/

class Solution
{
    private int $rows = 0;
    private int $cols = 0;
    private array $grid = [];
    private array $memo = [];

    public function uniquePathsWithObstacles(array $obstacleGrid): int
    {
        $this->grid = $obstacleGrid;
        $this->rows = count($obstacleGrid);
        $this->cols = count($obstacleGrid[0]);
        $this->memo = array_fill(0, $this->rows, []);
        $this->memo[$this->rows - 1][$this->cols - 1] = 1;

        return $this->dfs(0, 0);
    }


    public function dfs(int $r, int $c): int
    {
        if ($r === $this->rows || $c === $this->cols || $this->grid[$r][$c]) {
            return 0;
        }

        if (array_key_exists($c, $this->memo[$r])) {
            return $this->memo[$r][$c];
        }

        $this->memo[$r][$c] = $this->dfs($r + 1, $c) + $this->dfs($r, $c + 1);
        return $this->memo[$r][$c];
    }
}