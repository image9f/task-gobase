package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != mapping[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println("() is", isValid("()"))
	fmt.Println("()[]{} is", isValid("()[]{}"))
	fmt.Println("(] is", isValid("(]"))
	fmt.Println("([)] is", isValid("([)]"))
	fmt.Println("{[]} is", isValid("{[]}"))
}
