<?php


# https://leetcode.com/problems/kth-largest-element-in-an-array/
class Solution {
    public function findKthLargest($nums, $k) {
        $kLargest = [];

        foreach ($nums as $num) {
            if (count($kLargest) < $k) {
                $kLargest[] = $num;
            } else {
                $minValue = $this->findMinValue($kLargest);
                $minIndex = array_search($minValue, $kLargest);
                if ($num > $minValue) {
                    $kLargest[$minIndex] = $num;
                }
            }
        }

        return $this->findMinValue($kLargest);
    }

    private function findMinValue($array) {
        $minValue = PHP_INT_MAX;
        foreach ($array as $num) {
            if ($num < $minValue) {
                $minValue = $num;
            }
        }
        return $minValue;
    }
}
