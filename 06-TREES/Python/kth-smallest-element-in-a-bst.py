from typing import *
from TreeNode import TreeNode


# https://leetcode.com/problems/kth-smallest-element-in-a-bst/
class Solution:
    def __init__(self):
        self.inordered = []

    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        if root:
            self.kthSmallest(root.left, k)

            if len(self.inordered) < k:
                self.inordered.append(root.val)
            else:
                return self.inordered[k-1]

            self.kthSmallest(root.right, k)

        if len(self.inordered) >= k:
            return self.inordered[k - 1]


root = TreeNode.fromArray([5, 3, 6, 2, 4, None, None, 1])
root.print()
print(Solution().kthSmallest(root, 3))
