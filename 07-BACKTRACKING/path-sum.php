<?php

# https://leetcode.com/problems/path-sum/
class Solution
{
    public function hasPathSum(?TreeNode $root, int &$targetSum): bool
    {
        if (!$root) return false;

        $targetSum -= $root->val;

        $isLeafNode = (!$root->left) && (!$root->right);
        if ($isLeafNode && $targetSum == 0) {
            return true;
        }

        if ($this->hasPathSum($root->left, $targetSum) ||
            $this->hasPathSum($root->right, $targetSum)) {
            return true;
        }

        $targetSum += $root->val;
        return false;
    }
}