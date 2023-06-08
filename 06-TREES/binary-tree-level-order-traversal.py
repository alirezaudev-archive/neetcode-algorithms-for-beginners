from typing import *
from TreeNode import TreeNode
from collections import deque

# https://leetcode.com/problems/binary-tree-level-order-traversal/
class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        queue = deque()
        traversal = []

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
            traversal.append(tmp)

        return traversal





tree = TreeNode.fromArray([3, 9, 20, None, None, 15, 7])
Output = [[3], [9, 20], [15, 7]]

print(Solution().levelOrder(tree))
