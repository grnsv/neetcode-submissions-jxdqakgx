func checkInclusion(s1 string, s2 string) bool {
	counts := make(map[byte]int)
	for i := range len(s1) {
		counts[s1[i]]++
	}

	l := 0
	for r := range len(s2) {
		c, exists := counts[s2[r]]
		if c > 0 {
			counts[s2[r]]--

			if r - l + 1 == len(s1) {
				return true
			}

			continue
		} else if exists {
			for counts[s2[r]] == 0 {
				if _, ok := counts[s2[l]]; ok {
					counts[s2[l]]++
				}

				l++
			}

			counts[s2[r]]--

			if r - l + 1 == len(s1) {
				return true
			}
		} else {
			for l <= r {
				if _, ok := counts[s2[l]]; ok {
					counts[s2[l]]++
				}

				l++

				if s2[l-1] == s2[r] {
					break
				}
			}
		}
	}

	return false
}
