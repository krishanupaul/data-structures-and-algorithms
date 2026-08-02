// Problem: Valid Parentheses
// https://leetcode.com/problems/valid-parentheses/

package main

func isValid(s string) bool {
	var stack = make([]rune, 0, len(s))

	for _, currChar := range s {
		stackLen := len(stack)

		if stackLen > 0 {
			top := stack[stackLen-1]
			if (currChar == ')' && top == '(') || (currChar == ']' && top == '[') || (currChar == '}' && top == '{') {
				stack = stack[:stackLen-1]
				continue
			}
		}
		stack = append(stack, currChar)
	}

	return len(stack) == 0
}
