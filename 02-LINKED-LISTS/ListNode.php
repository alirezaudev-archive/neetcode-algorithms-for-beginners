<?php

class ListNode
{
    function __construct(
        public mixed $val = 0,
        public ?self $next = null,
        public ?self $prev = null
    )
    {
    }
}