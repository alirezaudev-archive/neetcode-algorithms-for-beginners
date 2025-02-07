<?php

# https://leetcode.com/problems/merge-k-sorted-lists/
class Solution
{
    public function mergeKLists($lists)
    {
        if (empty($lists)) {
            return null;
        }

        $head = new ListNode(0);
        $current = $head;

        foreach ($lists as $node) {
            $current->next = $node;
            while ($current->next) {
                $current = $current->next;
            }
        }

        return $this->sortList($head->next);
    }

    public function sortList($head)
    {
        if ($head === null || $head->next === null) {
            return $head;
        }

        $middle = $this->findMiddle($head);


        $leftHead = $head;
        $rightHead = $middle->next;
        $middle->next = null;


        $leftSorted = $this->sortList($leftHead);
        $rightSorted = $this->sortList($rightHead);

        return $this->merge($leftSorted, $rightSorted);
    }

    private function findMiddle($head) {
        $slow = $head;
        $fast = $head;

        while ($fast->next !== null && $fast->next->next !== null) {
            $slow = $slow->next;
            $fast = $fast->next->next;
        }

        return $slow;
    }

    private function merge($left, $right) {
        $dummy = new ListNode(0);
        $current = $dummy;

        while ($left !== null && $right !== null) {
            if ($left->val <= $right->val) {
                $current->next = $left;
                $left = $left->next;
            } else {
                $current->next = $right;
                $right = $right->next;
            }
            $current = $current->next;
        }

        $current->next = ($left !== null) ? $left : $right;

        return $dummy->next;
    }
}