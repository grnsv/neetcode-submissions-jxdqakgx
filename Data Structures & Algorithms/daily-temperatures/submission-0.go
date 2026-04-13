func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	type pair struct {
		t int
		i int
	}

	stack := make([]pair, 0, len(temperatures)/2)
	for i, t := range temperatures {
		for len(stack) > 0 {
			prev := stack[len(stack)-1]
			if prev.t >= t {
				break
			}

			res[prev.i] = i - prev.i
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, pair{t: t, i: i})
	}

	return res
}
