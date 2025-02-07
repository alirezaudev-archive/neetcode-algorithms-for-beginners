from typing import *
from TreeNode import TreeNode


# https://leetcode.com/problems/delete-node-in-a-bst/
class Solution:
    def deleteNode(self, root: TreeNode|None , key: int) -> TreeNode|None:
        if not root:
            return None

        if key > root.val:
            root.right = self.deleteNode(root.right, key)
        elif key < root.val:
            root.left = self.deleteNode(root.left, key)
        else:
            if not root.left:
                return root.right
            elif not root.right:
                return root.left
            else:
                minNode = self.minNode(root.right)
                root.val = minNode.val
                root.right = self.deleteNode(root.right, minNode.val)

        return root

    def minNode(self, root: TreeNode) -> TreeNode:
        if not root.left:
            return root

        return self.minNode(root.left)