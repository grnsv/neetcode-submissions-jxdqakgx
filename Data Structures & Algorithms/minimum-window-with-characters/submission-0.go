func minWindow(s string, t string) string {
	if len(t) == 0 {
		return t
	}

    counts := make(map[byte]int, len(t))
	for i := range len(t) {
		counts[t[i]]--
	}

	left, right := 0, 0
	bestLeft, bestRight := 0, 0
	found := 0
	for left <= right && right < len(s) {
		for found < len(t) && right < len(s) {
			count, exists := counts[s[right]]
			if exists {
				if count < 0 {
					found++
				}

				counts[s[right]] = count + 1
			}

			right++
		}

		if found < len(t) {
			break
		}

		for found == len(t) && left < right {
			count, exists := counts[s[left]]
			if exists {
				if count == 0 {
					break
				}

				counts[s[left]] = count - 1
			}

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
