<?php

# https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
class Solution
{
    public function lowestCommonAncestor(TreeNode $root, TreeNode $p, TreeNode $q): TreeNode
    {
        $p_path = [$root];
        $tempRoot = $root;
        while ($tempRoot and $tempRoot->val != $p->val) {
            $tempRoot = $tempRoot->val > $p->val ? $tempRoot->left : $tempRoot->right;
            $p_path[] = $tempRoot;
        }


        $q_path = [$root];
        $tempRoot = $root;
        while ($tempRoot and $tempRoot->val != $q->val) {
            $tempRoot = $tempRoot->val > $q->val ? $tempRoot->left : $tempRoot->right;
            $q_path[] = $tempRoot;
        }

        for ($i = count($q_path); $i >= 0; $i--) {
            if (in_array($q_path[$i], $p_path)) {
                return $p_path[$i];
            }
        }

        return $root;
    }
}