from typing import *


# https://leetcode.com/problems/number-of-students-unable-to-eat-lunch/
class Solution:
    def countStudents(self, students: List[int], sandwiches: List[int]) -> int:
        requeue = 0
        while students and requeue < len(students):
            student = students.pop(0)
            if student == sandwiches[0]:
                sandwiches.pop(0)
                requeue = 0
                continue
            students.append(student)
            requeue += 1

        return requeue
