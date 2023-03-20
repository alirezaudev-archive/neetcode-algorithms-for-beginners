<?php

# https://leetcode.com/problems/min-stack/
class MinStack
{
    private array $stack = [];
    private array $min_stack = [];

    public function push(int $val): void
    {
        $this->stack[] = $val;
        if (! empty($this->min_stack) && $val > $min = end($this->min_stack)) {
            $val = $min;
        }
        $this->min_stack[] = $val;
    }

    public function pop(): void
    {
        array_pop($this->min_stack);
        array_pop($this->stack);
    }

    public function top(): int
    {
        return end($this->stack);
    }

    public function getMin(): int
    {
        return end($this->min_stack);
    }
}
