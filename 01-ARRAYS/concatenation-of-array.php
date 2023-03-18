<?php

# https://leetcode.com/problems/concatenation-of-array/
function getConcatenation(array $nums): array
{
    return [...$nums, ...$nums];
}