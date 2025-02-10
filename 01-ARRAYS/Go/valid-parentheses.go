package main

// https://leetcode.com/problems/valid-parentheses/
func isValid(s string) bool {
	parentheses := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	parenthesesReverse := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	stack := make([]rune, len(s))
	i := 0
	for _, char := range s {
		if _, isOpen := parentheses[char]; isOpen {
			stack[i] = char
			i++
			continue
		}

		if i == 0 || stack[i-1] != parenthesesReverse[char] {
			return false
		}

		stack[i-1] = 0
		i--
	}

	return stack[0] == 0
}
