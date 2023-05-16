<?php

# https://leetcode.com/problems/search-in-a-binary-search-tree/
class Solution
{
    public function searchBTS(?TreeNode $root, int $val): ?TreeNode
    {
        if ($root === null)
            return null;

        if ($root->val == $val)
            return $root;

        if ($val < $root->val)
            return $this->searchBTS($root->left, $val);
        else
            return $this->searchBTS($root->right, $val);
    }

}
