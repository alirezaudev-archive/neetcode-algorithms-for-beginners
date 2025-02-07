<?php

# https://leetcode.com/problems/subsets/
class Solution {
    function subsets($nums) {
        $subs = [[]];

        foreach ($nums as $num) {
            $temp = [];
            foreach ($subs as $sub) {
                $temp[] = array_merge($sub, [$num]);
            }
            $subs = array_merge($subs, $temp);
        }

        return $subs;
    }
}
