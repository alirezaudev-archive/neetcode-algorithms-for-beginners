<?php

// https://leetcode.com/problems/design-hashmap/

class Node
{
    public function __construct(
        public int   $key,
        public int   $value,
        public ?Node $next = null,
    )
    {
    }
}

class MyHashMap
{
    private int $size = 1_000_003;
    private array $set = [];

    public function put(int $key, int $value): void
    {
        [$prev, $curr] = $this->getNode($key);
        if ($curr !== null) {
            $curr->value = $value;
            return;
        }

        $node = new Node($key, $value);
        if ($prev === null) {
            $this->set[$this->index($key)] = $node;
        } else {
            $prev->next = $node;
        }
    }

    public function remove(int $key): void
    {
        [$prev, $curr] = $this->getNode($key);
        if ($curr === null) {
            return;
        }

        if ($prev === null) {
            $this->set[$this->index($key)] = $curr->next;
        } else {
            $prev->next = $curr->next;
        }
    }

    public function get(int $key): int
    {
        $node = $this->getNode($key)[1];
        return $node ? $node->value : -1;
    }

    private function index(int $key): int
    {
        return $key % $this->size;
    }

    /**
     * @param int $key
     * @return Node[]
     */
    private function getNode(int $key): array
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