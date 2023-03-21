<?php

# https://leetcode.com/problems/reverse-linked-list/
class Solution {
    function reverseList(?ListNode $head): ?ListNode
    {
        $prev = null;
        $cur = $head;

        while ($cur !== null) {
            $next_node = $cur->next;
            $cur->next = $prev;
            $prev = $cur;
            $cur = $next_node;
        }

        return $prev;
    }
}
