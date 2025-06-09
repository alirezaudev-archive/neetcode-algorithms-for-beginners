<?php


class Solution
{
    function longestCommonSubsequence(string $text1, string $text2): int
    {
        if (strlen($text1) < strlen($text2)) {
            $tmp = $text1;
            $text1 = $text2;
            $text2 = $tmp;
        }

        $t1Len = strlen($text1);
        $t2Len = strlen($text2);

        $dp = array_fill(0, $t2Len, 0);

        for ($i = $t1Len - 1; $i >= 0; $i--) {
            $prev = 0;
            for ($j = $t2Len - 1; $j >= 0; $j--) {
                $tmp = $dp[$j];
                if ($text1[$i] === $text2[$j]) {
                    $dp[$j] = $prev + 1;
                } else {
                    $dp[$j] = max($dp[$j], $dp[$j + 1]);
                }
                $prev = $tmp;
            }
        }

        return $dp[0];
    }
}