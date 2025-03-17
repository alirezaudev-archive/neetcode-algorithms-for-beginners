<?php

// https://leetcode.com/problems/combination-sum/
class Solution
{
    private array $candidates = [];
    private int $target = 0;
    private array $res = [];

    /**
     * @param Integer[] $candidates
     * @param Integer $target
     * @return Integer[][]
     */
    function combinationSum(array $candidates, int $target): array
    {
        $this->target = $target;
        $this->candidates = $candidates;

        $curr = [];
        $this->dfs(0, $curr, 0);

        return $this->res;
    }

    private function dfs(int $index, array &$curr, int $total): void
    {
        if ($total == $this->target) {
            $this->res[] = $curr;
            return;
        }

        if ($total > $this->target || $index >= count($this->candidates)) {
            return;
        }

        $curr[] = $this->candidates[$index];
        $this->dfs($index, $curr, $total + $this->candidates[$index]);

        array_pop($curr);
        $this->dfs($index + 1, $curr, $total);
    }
}
