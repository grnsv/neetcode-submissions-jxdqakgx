func trap(height []int) int {
	res := 0
	if len(height) < 3 {
		return res
	}

	l, r := 0, len(height) - 1
	maxL, maxR := l, r
	for l < r {
		if height[maxL] < height[maxR] {
			l++
			if height[l] > height[maxL] {
				maxL = l
			}
			res += height[maxL] - height[l]
		} else {
			r--
			if height[r] > height[maxR] {
				maxR = r
			}
			res += height[maxR] - height[r]
		}
	}

	return res
}
