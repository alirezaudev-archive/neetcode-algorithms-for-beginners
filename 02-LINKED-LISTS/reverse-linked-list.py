from typing import *
import ListNode

# https://leetcode.com/problems/reverse-linked-list/
class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        prev = None
        cur = head

        while cur:
            next_node = cur.next
            cur.next = prev
            prev = cur
            cur = next_node

        return prev

