from TreeNode import TreeNode
from collections import deque


# https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        p_path = [root]
        tempRoot = root
        while tempRoot and tempRoot.val != p.val:
            tempRoot = tempRoot.left if tempRoot.val > p.val else tempRoot.right
            p_path.append(tempRoot)

        q_path = [root]
        tempRoot = root
        while tempRoot and tempRoot.val != q.val:
            tempRoot = tempRoot.left if tempRoot.val > q.val else tempRoot.right
            q_path.append(tempRoot)

        for i in range(len(q_path) - 1, -1, -1):
            if q_path[i] in p_path:
                return q_path[i]

        return root
