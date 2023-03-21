<?php

class ListNode
{
    public int $val = 0;
    public ?self $next = null;

    function __construct($val = 0, $next = null)
    {
        $this->val = $val;
        $this->next = $next;
    }
}