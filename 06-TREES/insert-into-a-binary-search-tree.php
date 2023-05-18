<?php

# https://leetcode.com/problems/insert-into-a-binary-search-tree/
class Solution
{
    public function insertIntoBST(?TreeNode $root, int $val): TreeNode
    {
        if (!$root)
            return new TreeNode($val);

        if ($val > $root->val)
            $root->right = $this->insertIntoBST($root->right, $val);
        elseif ($val < $root->val)
            $root->left = $this->insertIntoBST($root->left, $val);

        return $root;
    }
}