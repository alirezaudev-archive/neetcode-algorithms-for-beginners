<?php

# https://leetcode.com/problems/sort-an-array/
class InsertionSortingSolution
{
    /**
     * Insertion Sorting Algorithm
     */
    public function sortArray(array $nums): array
    {
        for ($i = 1; $i < count($nums); $i++) {
            $j = $i - 1;
            while ($j >= 0 && $nums[$j + 1] < $nums[$j]) {
                $tmp = $nums[$j];
                $nums[$j] = $nums[$j + 1];
                $nums[$j + 1] = $tmp;
                $j--;
            }

        }
        return $nums;
    }
}

class MergeSortingSolution
{
    /**
     * Merge Sorting Algorithm
     */
    public function sortArray(array $nums): array
    {
        return $this->mergeSort($nums, 0, count($nums) - 1);
    }

    private function mergeSort(array $nums, int $start, int $end): array
    {
        if ($end - $start + 1 <= 1) {
            return array_slice($nums, $start, $end - $start + 1);
        }

        $middle = (int)(($start + $end) / 2);

        $leftSorted = $this->mergeSort($nums, $start, $middle);
        $rightSorted = $this->mergeSort($nums, $middle + 1, $end);

        return $this->merge($leftSorted, $rightSorted);
    }

    private function merge(array $left, array $right): array
    {
        $merged = [];
        $i_left = $i_right = 0;

        while ($i_left < count($left) && $i_right < count($right)) {
            if ($left[$i_left] <= $right[$i_right]) {
                $merged[] = $left[$i_left];
                $i_left++;
            } else {
                $merged[] = $right[$i_right];
                $i_right++;
            }
        }

        while ($i_left < count($left)) {
            $merged[] = $left[$i_left];
            $i_left++;
        }

        while ($i_right < count($right)) {
            $merged[] = $right[$i_right];
            $i_right++;
        }

        return $merged;
    }
}

class QuickSortingSolution
{
    /**
     * Quick Sorting Algorithm
     */
    public function sortArray(array $nums): array
    {
        return $this->quickSort($nums, 0, count($nums) - 1);
    }

    private function quickSort(array &$nums, int $start, int $end): array
    {
        if ($end - $start + 1 <= 1) {
            return $nums;
        }

        $pivot = $nums[$end];
        $pointer = $start;

        for ($i = $start; $i < $end + 1; $i++) {
            if ($nums[$i] < $pivot) {
                $tmp = $nums[$pointer];
                $nums[$pointer] = $nums[$i];
                $nums[$i] = $tmp;
                $pointer++;
            }
        }

        $nums[$end] = $nums[$pointer];
        $nums[$pointer] = $pivot;

        $this->quickSort($nums, $start, $pointer - 1);
        $this->quickSort($nums,$pointer + 1, $end);

        return $nums;
    }
}


$nums = [2, 3, 4, 1, 6];
$insertionSorted = (new InsertionSortingSolution())->sortArray($nums);
$mergeSorted = (new MergeSortingSolution())->sortArray($nums);
$quickSorted = (new QuickSortingSolution())->sortArray($nums);

print_r($insertionSorted);
print_r($mergeSorted);
print_r($quickSorted);