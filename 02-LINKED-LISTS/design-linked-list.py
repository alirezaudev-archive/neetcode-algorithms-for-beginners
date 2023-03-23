from typing import *
from ListNode import ListNode


# https://leetcode.com/problems/design-linked-list/
class MyLinkedList:

    def __init__(self):
        self.head = ListNode()
        self.tail = None
        self.size = 0

    def get(self, index: int) -> int:
        if index < 0 or index >= self.size:
            return -1

        current = self.head.next
        for i in range(index):
            current = current.next

        return current.val

    def addAtHead(self, val: int) -> None:
        newNode = ListNode(val, self.head.next, None)
        if self.head.next:
            self.head.next.prev = newNode

        if self.tail is None:
            self.tail = newNode

        self.head.next = newNode
        self.size += 1

    def addAtTail(self, val: int) -> None:
        newNode = ListNode(val, None, self.tail)

        if self.tail:
            self.tail.next = newNode
        else:
            newNode.prev = self.head
            self.head.next = newNode

        self.tail = newNode
        self.size += 1

    def addAtIndex(self, index: int, val: int) -> None:
        if index < 0 or index > self.size:
            return

        if index == 0:
            self.addAtHead(val)
        elif index == self.size:
            self.addAtTail(val)
        else:
            current = self.head.next
            for i in range(index - 1):
                current = current.next

            newNode = ListNode(val, current.next, current)
            current.next.prev = newNode
            current.next = newNode
            self.size += 1

    def deleteAtIndex(self, index: int) -> None:
        if index < 0 or index >= self.size:
            return

        current = self.head.next
        for i in range(index):
            current = current.next

        if current.prev:
            current.prev.next = current.next
        else:
            self.head.next = current.next

        if current.next:
            current.next.prev = current.prev
        else:
            self.tail = current.prev

        self.size -= 1