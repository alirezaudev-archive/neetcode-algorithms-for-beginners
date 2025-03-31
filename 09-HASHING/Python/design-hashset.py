# https://leetcode.com/problems/design-hashset/

class Node:
    def __init__(self, key: int):
        self.key = key
        self.next = None


class MyHashSet:
    size = 10007

    def __init__(self):
        self.set = [None] * self.size

    def add(self, key: int) -> None:
        tail, curr = self._get(key)
        if curr is not None:
            return

        node = Node(key=key)
        if tail is None:
            self.set[self._index(key)] = node
        else:
            tail.next = node

    def remove(self, key: int) -> None:
        prev, curr = self._get(key)
        if curr is None:
            return

        if prev is None:
            self.set[self._index(key)] = curr.next
        else:
            prev.next = curr.next

    def contains(self, key: int) -> bool:
        return self._get(key)[1] is not None

    def _index(self, key: int) -> int:
        return key % self.size

    def _get(self, key: int):
        index = self._index(key)
        curr = self.set[index]
        prev = None
        while curr:
            if curr.key == key:
                return prev, curr
            prev = curr
            curr = curr.next
        return prev, None
