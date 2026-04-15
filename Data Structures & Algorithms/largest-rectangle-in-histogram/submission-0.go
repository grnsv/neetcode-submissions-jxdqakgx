func largestRectangleArea(heights []int) int {
	type pair struct {
		i int
		h int
	}

	res := 0
	stack := []pair{}
	for i, h := range heights {
		if h == 0 {
			stack = stack[:0]
		}

		res = max(res, h)
		stack = append(stack, pair{i: i, h: h})
		if len(stack) == 1 {
			continue
		}

		for j := len(stack)-2; j >= 0; j-- {
			prev := stack[j]
			res = max(res, min(h, prev.h) * (i - prev.i + 1))
			if prev.h > h {
				prev.h = h
				stack[j] = prev
				stack = stack[:j+1]
			}
		}
	}

	return res
}
