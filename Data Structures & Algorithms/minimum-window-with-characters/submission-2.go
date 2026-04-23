func minWindow(s string, t string) string {
	if len(t) == 0 {
		return t
	}

	counts := [128]int{}
	for i := range len(t) {
		counts[t[i]]--
	}

	l := 0
	bestL := 0
	bestLen := len(s) + 1
	found := 0

	for r := range len(s) {
		if counts[s[r]] < 0 {
			found++
		}
		counts[s[r]]++

		for found == len(t) {
			windowLen := r - l + 1
			if windowLen < bestLen {
				bestLen = windowLen
				bestL = l
			}

			counts[s[l]]--
			if counts[s[l]] < 0 {
				found--
			}
			l++
		}
	}

	if bestLen > len(s) {
		return string(s[:0])
	}

	return string(s[bestL : bestL+bestLen])
}
