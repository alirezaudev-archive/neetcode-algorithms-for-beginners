from typing import *


# https://leetcode.com/problems/merge-k-sorted-lists/
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode] | None:
        if not lists:
            return None

        head = ListNode(0)
        current = head

        for node in lists:
            current.next = node
            while current.next:
                current = current.next

        return self.sortList(head.next)

    def sortList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head or not head.next:
            return head

        middle = self.findMiddle(head)

        left_head = head
        right_head = middle.next
        middle.next = None

        left_sorted = self.sortList(left_head)
        right_sorted = self.sortList(right_head)

        return self.merge(left_sorted, right_sorted)

    def findMiddle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        slow = fast = head

        while fast.next and fast.next.next:
            slow = slow.next
            fast = fast.next.next

        return slow

    def merge(self, left: Optional[ListNode], right: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(0)
        current = dummy

        while left and right:
            if left.val <= right.val:
                current.next = left
                left = left.next
            else:
                current.next = right
                right = right.next
            current = current.next

        current.next = left if left else right

        return dummy.next
