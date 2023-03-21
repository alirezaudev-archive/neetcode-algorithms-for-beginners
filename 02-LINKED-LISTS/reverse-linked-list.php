<?php

# https://leetcode.com/problems/reverse-linked-list/
class ListNode
{
    public int $val = 0;
    public ?self $next = null;

    function __construct($val = 0, $next = null)
    {
        $this->val = $val;
        $this->next = $next;
    }
}

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
