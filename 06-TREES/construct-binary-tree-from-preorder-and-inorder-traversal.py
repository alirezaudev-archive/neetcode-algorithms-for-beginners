from typing import *
from TreeNode import TreeNode


# https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> TreeNode | None:
        if not preorder or not inorder:
            return None

        root_val = preorder.pop(0)
        root_index = inorder.index(root_val)

        root = TreeNode(root_val)
        root.left = self.buildTree(preorder, inorder[:root_index])
        root.right = self.buildTree(preorder, inorder[root_index + 1:])

        return root


preorder = [3, 9, 20, 15, 7]
inorder  = [9, 3, 15, 20, 7]
Output   = [3, 9, 20, None, None, 15, 7]

TreeNode.fromArray(Output).print()
print()
Solution().buildTree(preorder, inorder).print()