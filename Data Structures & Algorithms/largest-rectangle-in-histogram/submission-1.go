func largestRectangleArea(heights []int) int {
	maxArea := 0
	n := len(heights)
	stack := []int{}
	for i := range n + 1 {
		for len(stack) > 0 && (i == n || heights[stack[len(stack)-1]] >= heights[i]) {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			w := i
			if len(stack) > 0 {
				w = w - stack[len(stack)-1] - 1
			}

			maxArea = max(maxArea, h * w)
		}

		stack = append(stack, i)
	}

	return maxArea
}
