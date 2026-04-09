func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}
	stack := []int{}
	for _, token := range tokens {
		var a int
		switch token {
		case "+", "-", "*", "/":
			var b int
			a, b = stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				a += b
			case "-":
				a -= b
			case "*":
				a *= b
			case "/":
				a /= b
			}
		default:
			a, _ = strconv.Atoi(token)
		}
		stack = append(stack, a)
	}
	return stack[0]
}
