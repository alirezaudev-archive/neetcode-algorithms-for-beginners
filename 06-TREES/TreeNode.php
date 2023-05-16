<?php

class TreeNode
{
     function __construct(
         public mixed $val = null,
         public ?self $left = null,
         public ?self $right = null,
     )
     {
     }
}
