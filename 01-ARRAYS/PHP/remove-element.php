<?php

# https://leetcode.com/problems/remove-element/
function removeElement(array &$nums, int $val): int
{
    $k = 0;
    foreach ($nums as $num)
        if ($num != $val) {
            $nums[$k++] = $num;
        }
    return $k;
}

