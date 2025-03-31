<?php

// https://leetcode.com/problems/lru-cache/
final class Node
{
    public function __construct(
        public int   $k,
        public int   $v,
        public ?Node $next = null,
        public ?Node $prev = null,
    )
    {
    }
}

class LRUCache
{
    /**
     * @var Node[]
     */
    private array $data = [];

    private ?Node $head = null;

    private ?Node $tail = null;

    function __construct(private readonly int $capacity)
    {
    }

    public function get(int $key): int
    {
        if ($node = $this->data[$key] ?? false) {
            $this->moveToTail($node);
            return $node->v;
        }

        return -1;
    }

    public function put(int $key, int $value): void
    {
        if ($node = $this->data[$key] ?? false) {
            $node->v = $value;
            $this->moveToTail($node);
            return;
        }

        $node = new Node(k: $key, v: $value, prev: $this->tail);
        $this->data[$key] = $node;
        if ($this->head === null) {
            $this->head = $node;
            $this->tail = $node;
            return;
        }

        $this->tail->next = $node;
        $this->tail = $node;
        if (count($this->data) > $this->capacity) {
            $oldHead = $this->head;
            $this->head = $this->head->next;
            $this->head->prev = null;
            unset($this->data[$oldHead->k]);
        }
    }

    private function moveToTail(Node $node): void
    {
        if ($node === $this->tail) {
            return;
        }

        if ($this->head === $node) {
            $this->head = $this->head->next;
        } else {
            $node->next->prev = $node->prev;
            $node->prev->next = $node->next;
        }

        $node->next = null;
        $node->prev = $this->tail;
        $this->tail->next = $node;
        $this->tail = $node;
    }
}