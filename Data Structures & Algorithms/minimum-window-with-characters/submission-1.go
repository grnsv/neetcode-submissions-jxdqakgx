func minWindow(s string, t string) string {
	if len(t) == 0 {
		return t
	}

    counts := [128]int{}
	for i := range len(t) {
		counts[t[i]]--
	}

	left, right := 0, 0
	bestLeft, bestRight := 0, 0
	found := 0
	for left <= right && right < len(s) {
		for found < len(t) && right < len(s) {
			if counts[s[right]] < 0 {
				found++
			}

			counts[s[right]]++
			right++
		}

		if found < len(t) {
			break
		}

		for left < right {
			if counts[s[left]] == 0 {
				break
			}

			counts[s[left]]--
			left++
		}

		if left != right && right - left < bestRight - bestLeft || bestRight == bestLeft {
			bestLeft, bestRight = left, right
		}

		found--
		counts[s[left]]--
		left++
	}

	return string(s[bestLeft:bestRight])
}
