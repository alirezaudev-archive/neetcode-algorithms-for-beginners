from typing import *

# https://leetcode.com/problems/reverse-linked-list/
class Solution:
    def reverseList(self, head: ListNode):
        if not head or not head.next:
            return head

        new_head = self.reverseList(head.next)

        head.next.next = head
        head.next = None

        return new_head

