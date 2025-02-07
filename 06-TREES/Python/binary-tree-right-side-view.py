from typing import *
from TreeNode import TreeNode
from collections import deque


# https://leetcode.com/problems/binary-tree-right-side-view/
class Solution:
    def rightSideView(self, root: Optional[TreeNode]) -> List[int]:
        queue = deque()
        rightSide = []

        if root:
            queue.append(root)

        while len(queue) > 0:
            tmp = []
            for i in range(len(queue)):
                curr = queue.popleft()
                tmp.append(curr.val)
                if curr.left:
                    queue.append(curr.left)
                if curr.right:
                    queue.append(curr.right)
            rightSide.append(tmp[-1])

        return rightSide


tree = TreeNode.fromArray([1, 2, 3, 4])
tree.print()
print()
print(Solution().rightSideView(tree))
