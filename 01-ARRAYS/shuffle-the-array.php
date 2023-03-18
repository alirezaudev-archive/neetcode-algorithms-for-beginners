<?php

# https://leetcode.com/problems/shuffle-the-array/
function shuffleArray(array &$nums, int $n): array
{
    $array = [];
    for ($i = 0; $i < $n; $i++) {
        $array[] = $nums[$i];
        $array[] = $nums[$i + $n];
    }
    return $array;
}