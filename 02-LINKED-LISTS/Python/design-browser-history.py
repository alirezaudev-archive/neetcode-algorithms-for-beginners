from typing import *
from ListNode import ListNode

# https://leetcode.com/problems/design-browser-history/
class BrowserHistory:
    def __init__(self, homepage: str):
        self.current = ListNode(homepage)

    def visit(self, url: str) -> None:
        visitedNode = ListNode(url)
        visitedNode.prev = self.current
        self.current.next,  = visitedNode
        self.current = visitedNode

    def back(self, steps: int) -> str:
        while steps > 0 and self.current.prev:
            self.current = self.current.prev
            steps -= 1
        return self.current.val

    def forward(self, steps: int) -> str:
        while steps > 0 and self.current.next:
            self.current = self.current.next
            steps -= 1
        return self.current.val


