<?php


# https://leetcode.com/problems/binary-tree-level-order-traversal/
class Solution
{
    public function levelOrder(?TreeNode $root): array
    {
        $queue = [];
        $traversal = [];

        if (!$root) return [];

        $queue[] = $root;

        while (count($queue) > 0) {
            $size = count($queue);
            $tmp = [];

            for ($i = 0; $i < $size; $i++) {
                $curr = array_shift($queue);
                $tmp[] = $curr->val;

                if ($curr->left) {
                    $queue[] = $curr->left;
                }
                if ($curr->right) {
                    $queue[] = $curr->right;
                }
            }

            $traversal[] = $tmp;
        }

        return $traversal;
    }
}
