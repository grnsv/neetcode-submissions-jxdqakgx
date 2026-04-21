func characterReplacement(s string, k int) int {
	res := 0

	freqs := make(map[byte]int)
	maxFreq := 0

	l := 0
	for r := range s {
		freqs[s[r]]++
		maxFreq = max(maxFreq, freqs[s[r]])

		for r - l + 1 - maxFreq > k {
			freqs[s[l]]--
			l++
		}

		res = max(res, r - l + 1)
	}

	return res
}
