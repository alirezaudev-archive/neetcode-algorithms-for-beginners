<?php
use TreeNode;

# https://leetcode.com/problems/binary-tree-inorder-traversal/
class Solution
{
    private array $inordered = [];

    public function inorderTraversal(?TreeNode $root): array
    {
        if ($root) {
            $this->inorderTraversal($root->left);
            $this->inordered[] = $root->val;
            $this->inorderTraversal($root->right);
        }

        return $this->inordered;
    }
}