<?php
require '../vendor/autoload.php';

# https://leetcode.com/problems/concatenation-of-array/
function getConcatination(array $nums): array
{
    return [...$nums, ...$nums];
}