def reverseString(string: str):
    if len(string) <= 1:
        return string

    return string[-1] + reverseString(string[:-1])


def isPalindrome(string: str):
    if len(string) <= 1:
        return True

    if string[0] == string[-1]:
        return isPalindrome(string[1:-1])

    return False


def findBinary(n: int):
    mod = n % 2
    if n <= 1:
        return mod

    return str(findBinary(n // 2)) + str(n % 2)


def recursiveSummation(num: int):
    if num <= 1:
        return num

    return num + recursiveSummation(num - 1)


