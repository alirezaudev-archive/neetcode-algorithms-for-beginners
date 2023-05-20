from typing import *
from TreeNode import TreeNode

# https://leetcode.com/problems/binary-tree-inorder-traversal/
class Solution:
    def __init__(self):
        self.inordered = []

    def inorderTraversal(self, root: TreeNode | None) -> List[int]:
        if root:
            self.inorderTraversal(root.left)
            self.inordered.append(root.val)
            self.inorderTraversal(root.right)

        return self.inordered
