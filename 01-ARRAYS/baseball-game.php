<?php

# https://leetcode.com/problems/baseball-game/
function calPoints(array $operations): int
{
    $stack = [];

    foreach ($operations as $i) {
        if ($i == '+') {
            $stack[] = array_sum(array_slice($stack, -2));
        } elseif ($i == 'D') {
            $stack[] = end($stack) * 2;
        } elseif ($i == 'C') {
             array_pop($stack);
        } else{
            $stack[] = $i;
        }
    }

    return array_sum($stack);
}