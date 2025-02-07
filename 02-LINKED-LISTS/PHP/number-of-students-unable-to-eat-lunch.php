<?php

# https://leetcode.com/problems/number-of-students-unable-to-eat-lunch/
class Solution
{
    function countStudents(array $students, array $sandwiches): int
    {
        $requeue = 0;
        while ($students && $requeue < count($students)) {
            $student = array_shift($students);
            if ($student == $sandwiches[0]) {
                array_shift($sandwiches);
                $requeue = 0;
                continue;
            }
            $students[] = $student;
            $requeue++;
        }

        return $requeue;
    }
}
