func isValid(s string) bool {
    stack := make([]rune, 0, len(s))
	for _, r := range s {
		if r == ')' {
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		} else if r == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		} else if r == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, r)
		}
	}
	return len(stack) == 0
}
