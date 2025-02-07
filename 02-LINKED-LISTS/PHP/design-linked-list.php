<?php
use ListNode as MyListNode;

# https://leetcode.com/problems/design-linked-list/
class MyLinkedList {
    public ?MyListNode $head;
    public ?MyListNode $tail;
    public int $size;

    public function __construct() {
        $this->head = new MyListNode();
        $this->tail = null;
        $this->size = 0;
    }

    public function get(int $index): int
    {
        if ($index < 0 || $index >= $this->size) {
            return -1;
        }

        $current = $this->head->next;
        for ($i = 0; $i < $index; $i++) {
            $current = $current->next;
        }

        return $current->val;
    }

    public function addAtHead(int $val): void
    {
        $newNode = new MyListNode($val, $this->head->next, null);
        if ($this->head->next) {
            $this->head->next->prev = $newNode;
        }

        if ($this->tail === null) {
            $this->tail = $newNode;
        }

        $this->head->next = $newNode;
        $this->size++;
    }

    public function addAtTail(int $val): void
    {
        $newNode = new MyListNode($val, null, $this->tail);

        if ($this->tail) {
            $this->tail->next = $newNode;
        } else {
            $newNode->prev = $this->head;
            $this->head->next = $newNode;
        }

        $this->tail = $newNode;
        $this->size++;
    }

    public function addAtIndex(int $index, int $val): void
    {
        if ($index < 0 || $index > $this->size) {
            return;
        }

        if ($index === 0) {
            $this->addAtHead($val);
        } elseif ($index === $this->size) {
            $this->addAtTail($val);
        } else {
            $current = $this->head->next;
            for ($i = 0; $i < $index - 1; $i++) {
                $current = $current->next;
            }

            $newNode = new MyListNode($val, $current->next, $current);
            $current->next->prev = $newNode;
            $current->next = $newNode;
            $this->size++;
        }
    }

    public function deleteAtIndex(int $index): void
    {
        if ($index < 0 || $index >= $this->size) {
            return;
        }

        $current = $this->head->next;
        for ($i = 0; $i < $index; $i++) {
            $current = $current->next;
        }

        if ($current->prev) {
            $current->prev->next = $current->next;
        } else {
            $this->head->next = $current->next;
        }

        if ($current->next) {
            $current->next->prev = $current->prev;
        } else {
            $this->tail = $current->prev;
        }

        $this->size--;
    }
}
