<?php

/**
 * https://leetcode.com/problems/remove-duplicates-from-sorted-array/
 *
 * @param array<int> $nums
 * @return int
 */
function removeDuplicates(array &$nums): int
{
    $uniquesLength = 1;
    for ($i = 1; $i < count($nums); $i++)
        if ($nums[$i - 1] != $nums[$i]) {
            $nums[$uniquesLength] = $nums[$i];
            $uniquesLength++;
        }

    return $uniquesLength;
}
