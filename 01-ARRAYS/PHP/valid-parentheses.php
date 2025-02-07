<?php

# https://leetcode.com/problems/valid-parentheses/
function isValid(string $s): bool
{
    $validBrackets = [
        '{' => '}',
        '[' => ']',
        '(' => ')',
    ];

    $openBracketsStack = [];
    foreach (str_split($s) as $c) {
        if (array_key_exists($c, $validBrackets)) {
            $openBracketsStack[] = $c;
            continue;
        } elseif (! empty($openBracketsStack)) {
            if ($c != $validBrackets[end($openBracketsStack)])
                return false;
            array_pop($openBracketsStack);
            continue;
        }
        return false;
    }

    return empty($openBracketsStack);
}