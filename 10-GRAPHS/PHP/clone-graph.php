<?php

// https://leetcode.com/problems/clone-graph/

class Node
{
    function __construct(
        public ?int  $val = 0,
        public array $neighbors = []
    )
    {
    }
}


class Solution
{
    private array $visited = [];

    function cloneGraph(?Node $node): ?Node
    {
        if (!$node) {
            return null;
        }

        return $this->dfs($node);
    }

    private function dfs(Node $node): Node
    {
        if ($clone = $this->visited[$node->val] ?? null) {
            return $clone;
        }

        $clone = new Node(val: $node->val);
        $this->visited[$node->val] = $clone;

        foreach ($node->neighbors as $neighbor) {
            $clone->neighbors[] = $this->dfs($neighbor);
        }

        return $clone;
    }
}
