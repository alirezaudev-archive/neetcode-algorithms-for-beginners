<?php

class ListNode
{
    function __construct(
        public int $val = 0,
        public ?self $next = null,
        public ?self $prev = null
    )
    {
    }
}