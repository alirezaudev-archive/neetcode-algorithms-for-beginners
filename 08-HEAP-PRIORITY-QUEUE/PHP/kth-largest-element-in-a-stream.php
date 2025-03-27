<?php

// https://leetcode.com/problems/kth-largest-element-in-a-stream/

class KthLargest
{
    private array $heap = [];

    function __construct(private readonly int $k, array $nums)
    {
        foreach ($nums as $num) {
            $this->add($num);
        }
    }

    function add(int $val): int
    {
        $length = count($this->heap);
        if ($length < $this->k) {
            $this->heap[] = $val;
            $i = $length;
            while ($i > 0) {
                $parent = floor(($i - 1) / 2);
                if ($this->heap[$parent] <= $this->heap[$i]) {
                    break;
                }

                $tmp = $this->heap[$i];
                $this->heap[$i] = $this->heap[$parent];
                $this->heap[$parent] = $tmp;
                $i = $parent;
            }


        } else if ($val > $this->heap[0]) {
            $this->heap[0] = $val;
            $i = 0;
            $n = $this->k;

            while (true) {
                $smallest = $i;
                $left = 2 * $i + 1;
                $right = 2 * $i + 2;

                if ($left < $n && $this->heap[$left] < $this->heap[$smallest]) {
                    $smallest = $left;
                }

                if ($right < $n && $this->heap[$right] < $this->heap[$smallest]) {
                    $smallest = $right;
                }

                if ($smallest == $i) {
                    break;
                }

                $tmp = $this->heap[$i];
                $this->heap[$i] = $this->heap[$smallest];
                $this->heap[$smallest] = $tmp;

                $i = $smallest;
            }
        }

        return $this->heap[0];
    }
}