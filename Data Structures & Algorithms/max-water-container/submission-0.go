func maxArea(heights []int) int {
	res := 0
	l := 0
	r := len(heights) - 1
	for l < r {
		left := heights[l]
		right := heights[r]
		area := min(left, right) * (r - l)
		res = max(res, area)
		if left < right {
			for l < r {
				l++
				if heights[l] > left {
					break
				}
			}
		} else {
			for l < r {
				r--
				if heights[r] > right {
					break
				}
			}
		}
	}

	return res
}
