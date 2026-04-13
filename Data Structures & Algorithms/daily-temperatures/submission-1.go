func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([]int, 0, len(temperatures)/2)
	for i, t := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < t {
			prev := stack[len(stack)-1]
			res[prev] = i - prev
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return res
}
