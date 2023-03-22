from typing import *
import ListNode


# https://leetcode.com/problems/merge-two-sorted-lists/
class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        cur1 = list1
        cur2 = list2

        mergedHead = ListNode()
        current = mergedHead

        while cur1 and cur2:
            if cur2.val < cur1.val:
                current.next = cur2
                cur2 = cur2.next
            else:
                current.next = cur1
                cur1 = cur1.next
            current = current.next

        if cur1:
            current.next = cur1
        elif cur2:
            current.next = cur2

        return mergedHead.next