
#https://leetcode.com/problems/lru-cache/
class Node:
    def __init__(self, k: int, v: int, next: 'Node' = None, prev: 'Node' = None):
        self.k = k
        self.v = v
        self.next = next
        self.prev = prev

class LRUCache:
    def __init__(self, capacity: int):
        self.capacity = capacity
        self.data = {}
        self.head = None
        self.tail = None

    def get(self, key: int) -> int:
        if key in self.data:
            node = self.data[key]
            self._move_to_tail(node)
            return node.v
        return -1

    def put(self, key: int, value: int) -> None:
        if key in self.data:
            node = self.data[key]
            node.v = value
            self._move_to_tail(node)
            return

        node = Node(k=key, v=value, prev=self.tail)
        self.data[key] = node

        if self.head is None:
            self.head = node
            self.tail = node
            return

        self.tail.next = node
        self.tail = node

        if len(self.data) > self.capacity:
            old_head = self.head
            self.head = self.head.next
            self.head.prev = None
            del self.data[old_head.k]

    def _move_to_tail(self, node: Node) -> None:
        if node == self.tail:
            return

        if self.head == node:
            self.head = self.head.next
        else:
            node.next.prev = node.prev
            node.prev.next = node.next

        node.next = None
        node.prev = self.tail
        self.tail.next = node
        self.tail = node