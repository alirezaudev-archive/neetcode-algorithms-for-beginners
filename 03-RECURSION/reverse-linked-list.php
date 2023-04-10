<?php

# https://leetcode.com/problems/reverse-linked-list/
class Solution
{
    public function reverseList(?ListNode $head): ?ListNode
    {
        if (!$head or !$head->next) {
            return $head;
        }

        $new_head = $this->reverseList($head->next);

        $head->next->next = $head;
        $head->next = null;

        return $new_head;
    }
}
