<?php

# https://leetcode.com/problems/delete-node-in-a-bst/
class Solution
{
    public function deleteNode(?TreeNode $root, int $key): ?TreeNode
    {
        if (!$root)
            return null;

        if ($key > $root->val) {
            $root->right = $this->deleteNode($root->right, $key);
        } elseif ($key < $root->val) {
            $root->left = $this->deleteNode($root->left, $key);
        } else {
            if (!$root->left) {
                return $root->right;
            } elseif (!$root->right) {
                return $root->left;
            } else {
                $minNode = $this->minNode($root->right);
                $root->val = $minNode->val;
                $root->right = $this->deleteNode($root->right, $minNode->val);
            }
        }

        return $root;
    }

    private function minNode(TreeNode $root): TreeNode
    {
        while ($root and $root->left)
            $root = $root->left;

        return $root;
    }
}