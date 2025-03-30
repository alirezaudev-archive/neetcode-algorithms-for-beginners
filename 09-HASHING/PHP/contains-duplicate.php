<?php

// https://leetcode.com/problems/contains-duplicate/

class Solution
{
    function containsDuplicate(array $nums): bool
    {
        $h = [];
        foreach ($nums as $num) {
            if (key_exists($num, $h)) {
                return false;
            }
            $h[$num] = 1;
        }

        return true;
    }
}