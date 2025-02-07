<?php


# https://leetcode.com/problems/implement-stack-using-queues/
class MyStack {
    private $queue;

    public function __construct() {
        $this->queue = new SplQueue();
    }

    public function push($x) {
        $this->queue->enqueue($x);
        $count = $this->queue->count();

        for ($i = 0; $i < $count - 1; $i++) {
            $this->queue->enqueue($this->queue->dequeue());
        }
    }

    public function pop() {
        return $this->queue->dequeue();
    }

    public function top() {
        return $this->queue->bottom();
    }

    public function empty() {
        return $this->queue->isEmpty();
    }
}