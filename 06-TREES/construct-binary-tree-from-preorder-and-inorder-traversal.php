<?php


# https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
class Solution
{
    function buildTree(array &$preorder, array $inorder): ?TreeNode
    {
        if (empty($preorder) || empty($inorder)) {
            return null;
        }

        $root_val = array_shift($preorder);
        $root_index = array_search($root_val, $inorder);

        $root = new TreeNode($root_val);
        $root->left = $this->buildTree($preorder, array_slice($inorder, 0, $root_index));
        $root->right = $this->buildTree($preorder, array_slice($inorder, $root_index + 1));

        return $root;
    }
}


$preorder = [3, 4, 6, 1, 2, 7, 5, 8, 10, 9];
$inorder  = [1, 6, 2, 4, 7, 3, 8, 10, 5, 9];
$output   = [3, 4, 5, 6, 7, 8, 9, 1, 2, null, null, null, 10];

print_r((new Solution())->buildTree($preorder, $inorder));