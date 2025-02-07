from typing import *
from TreeNode import TreeNode


# https://leetcode.com/problems/path-sum/
class Solution:
    def hasPathSum(self, root: Optional[TreeNode], targetSum: int) -> bool:
        if not root:
            return False

        targetSum -= root.val

        isLeafNode = not root.left and not root.right
        if isLeafNode and targetSum == 0:
            return True

        if self.hasPathSum(root.left, targetSum) or \
                self.hasPathSum(root.right, targetSum):
            return True

        targetSum += root.val
        return False


print(Solution().hasPathSum(
    TreeNode.fromArray([1, 2]).print(),
    1
))
