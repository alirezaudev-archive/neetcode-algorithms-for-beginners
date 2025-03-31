# https://leetcode.com/problems/design-hashmap/

class Node:
    def __init__(self, key: int, value: int):
        self.key = key
        self.value = value
        self.next = None


class MyHashMap:
    size = 1_000_003

    def __init__(self):
        self.set = [None] * self.size

    def put(self, key: int, value: int) -> None:
        tail, curr = self._get(key)
        if curr is not None:
            curr.value = value
            return

        node = Node(key=key, value=value)
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

    def get(self, key: int) -> int:
        node = self._get(key)[1]
        return -1 if node is None else node.value

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
