<?php


# https://leetcode.com/problems/binary-tree-right-side-view/
class Solution
{
    function rightSideView(?TreeNode $root)
    {
        $queue = [];
        $rightSide = [];

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

            $rightSide[] = end($tmp);
        }

        return $rightSide;
    }

    /**
     * The solution below is so good, whoever wrote it.
     */

    private array $res = [];

    function rightSideView_v2($root) {
        $this->travel($root, 0);
        return $this->res;
    }

    function travel($node, $level) {
        if (is_null($node)) {
            return;
        }
        $this->res[$level] = $node->val;
        if ($left = $node->left) {
            $this->travel($left, $level + 1);
        }
        if ($right = $node->right) {
            $this->travel($right, $level + 1);
        }
    }
}