<?php


# https://leetcode.com/problems/merge-two-sorted-lists/
class Solution
{
    public function mergeTwoLists(?ListNode $list1, ?ListNode $list2): ?ListNode
    {
        $cur1 = $list1;
        $cur2 = $list2;

        $mergedHead = new ListNode();
        $current = $mergedHead;

        while ($cur1 && $cur2) {
            if ($cur2->val < $cur1->val) {
                $current->next = $cur2;
                $cur2 = $cur2->next;
            } else {
                $current->next = $cur1;
                $cur1 = $cur1->next;
            }
            $current = $current->next;
        }

        if ($cur1)
            $current->next = $cur1;
        elseif ($cur2)
            $current->next = $cur2;

        return $mergedHead->next;
    }
}