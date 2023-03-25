<?php

# https://leetcode.com/problems/design-browser-history/
class BrowserHistory
{
    private ListNode $current;

    public function __construct(string $homepage)
    {
        $this->current = new ListNode($homepage);
    }

    public function visit(string $url): void
    {
        $visitedNode = new ListNode($url);
        $visitedNode->prev = $this->current;
        $this->current->next = $visitedNode;
        $this->current = $visitedNode;
    }

    public function back(int $steps): string
    {
        while ($steps > 0 && $this->current->prev) {
            $this->current = $this->current->prev;
            $steps--;
        }
        return $this->current->val;
    }

    public function forward(int $steps): string
    {
        while ($steps > 0 && $this->current->next) {
            $this->current = $this->current->next;
            $steps--;
        }
        return $this->current->val;
    }
}