from typing import *


# https://leetcode.com/problems/baseball-game/
def calPoints(operations: List[str]) -> int:
    stack = []

    helpers = {
        '+': lambda: stack.append(stack[-1] + stack[-2]),
        'D': lambda: stack.append(stack[-1] * 2),
        'C': lambda: stack.pop(),
    }

    [helpers.get(i, lambda: stack.append(int(i)))() for i in operations]

    return sum(stack)

