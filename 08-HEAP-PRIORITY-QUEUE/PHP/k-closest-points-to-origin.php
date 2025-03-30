<?php

class Solution
{
    private int $k;
    private array $points = [];

    function kClosest(array $points, int $k): array
    {
        $this->k = $k;

        foreach ($points as $point) {
            $this->add($point);
        }

        return array_map(fn($point) => array_slice($point, 0, 2), $this->points);
    }

    public function add(array $point): void
    {
        $point[] = $this->calcDistance($point);
        $length = count($this->points);

        if ($length < $this->k) {
            $this->points[] = $point;
            $i = $length;
            while ($i > 0 && $this->points[intdiv($i - 1, 2)][2] < $this->points[$i][2]) {
                $parent = intdiv($i - 1, 2);
                [$this->points[$i], $this->points[$parent]] = [$this->points[$parent], $this->points[$i]];
                $i = $parent;
            }
        } elseif ($point[2] < $this->points[0][2]) {
            $this->points[0] = $point;
            $i = 0;
            while (true) {
                $largest = $i;
                $left = 2 * $i + 1;
                $right = 2 * $i + 2;
                if ($left < $this->k && $this->points[$largest][2] < $this->points[$left][2]) {
                    $largest = $left;
                }
                if ($right < $this->k && $this->points[$largest][2] < $this->points[$right][2]) {
                    $largest = $right;
                }
                if ($largest === $i) {
                    break;
                }
                [$this->points[$i], $this->points[$largest]] = [$this->points[$largest], $this->points[$i]];
                $i = $largest;
            }
        }
    }

    private function calcDistance(array $point): int
    {
        return $point[0] ** 2 + $point[1] ** 2;
    }
}