<?php

// https://leetcode.com/problems/design-hashset/

class Node
{
    public function __construct(
        public int   $key,
        public ?Node $next = null,
    )
    {
    }
}

class MyHashSet
{
    private int $size = 10007;
    private array $set = [];

    public function add(int $key): void
    {
        [$prev, $curr] = $this->get($key);
        if ($curr !== null) {
            return;
        }

        $node = new Node($key);
        if ($prev === null) {
            $this->set[$this->index($key)] = $node;
        } else {
            $prev->next = $node;
        }
    }

    public function remove(int $key): void
    {
        [$prev, $curr] = $this->get($key);
        if ($curr === null) {
            return;
        }

        if ($prev === null) {
            $this->set[$this->index($key)] = $curr->next;
        } else {
            $prev->next = $curr->next;
        }
    }

    public function contains(int $key): bool
    {
        return $this->get($key)[1] !== null;
    }

    private function index(int $key): int
    {
        return $key % $this->size;
    }

    private function get(int $key): array
    {
        $index = $this->index($key);
        $curr = $this->set[$index];
        $prev = null;

        while ($curr !== null) {
            if ($curr->key === $key) {
                return [$prev, $curr];
            }
            $prev = $curr;
            $curr = $curr->next;
        }

        return [$prev, null];
    }
}