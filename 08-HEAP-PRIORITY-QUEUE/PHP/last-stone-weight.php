<?php

// https://leetcode.com/problems/last-stone-weight/

class Solution
{
    private array $heap = [];

    public function lastStoneWeight(array $stones): int
    {
        foreach ($stones as $stone) {
            $this->push($stone);
        }

        while (count($this->heap) > 2) {
            $y = $this->pop();
            $x = $this->pop();
            if ($y != $x) {
                $this->push($y - $x);
            }
        }

        $len = count($this->heap);
        if ($len == 2) {
            return max($this->heap) - min($this->heap);
        }

        return $this->heap[0] ?? 0;
    }

    private function push(int $stone): void
    {
        $this->heap[] = $stone;
        $i = count($this->heap) - 1;
        while ($i > 0 && $this->heap[floor(($i - 1) / 2)] < $this->heap[$i]) {
            $parent = floor(($i - 1) / 2);
            $tmp = $this->heap[$parent];
            $this->heap[$parent] = $this->heap[$i];
            $this->heap[$i] = $tmp;
            $i = $parent;
        }
    }

    private function pop(): int
    {
        if (count($this->heap) < 1) {
            return -1;
        }

        $res = $this->heap[0];
        $this->heap[0] = array_pop($this->heap);
        $i = 0;
        $m = $this->maxChild($i);
        while ($m < count($this->heap) && $this->heap[$i] < $this->heap[$m]) {
            $tmp = $this->heap[$m];
            $this->heap[$m] = $this->heap[$i];
            $this->heap[$i] = $tmp;
            $i = $m;
            $m = $this->maxChild($i);
        }

        return $res;
    }

    private function maxChild(int $i): int
    {
        $len = count($this->heap);
        $left = 2 * $i + 1;
        $right = 2 * $i + 2;
        if ($right < $len && $this->heap[$right] > $this->heap[$left]) {
            return $right;
        }
        if ($left < $len) {
            return $left;
        }
        return $i;
    }
}