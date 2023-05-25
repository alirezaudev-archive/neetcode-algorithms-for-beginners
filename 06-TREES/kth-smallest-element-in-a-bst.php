<?php

# https://leetcode.com/problems/kth-smallest-element-in-a-bst/
class Solution
{
    private array $inordered = [];

    public function kthSmallest(?TreeNode $root, int $k): ?int
    {
        if ($root) {
            $this->kthSmallest($root->left, $k);

            if (count($this->inordered) < $k)
                $this->inordered[] = $root->val;
            else
                return $this->inordered[$k - 1];

            $this->kthSmallest($root->right, $k);
        }

        return $this->inordered[$k - 1] ?? null;
    }
}

