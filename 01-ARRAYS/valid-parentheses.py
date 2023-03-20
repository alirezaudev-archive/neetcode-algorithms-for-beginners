from typing import *


# https://leetcode.com/problems/valid-parentheses/
def isValid(s: str) -> bool:
    validBrackets = {'{': '}', '[': ']', '(': ')'}

    openBracketsStack = []
    for c in s:
        if c in validBrackets:
            openBracketsStack.append(c)
            continue
        elif openBracketsStack:
            if c != validBrackets[openBracketsStack[-1]]:
                return False

            openBracketsStack.pop()
            continue

        return False

    return not openBracketsStack
