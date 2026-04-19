func lengthOfLongestSubstring(s string) int {
	res := 0
	m := map[byte]int{}
	l, r := 0, 0
	for r < len(s) {
		if i, ok := m[s[r]]; ok {
			res = max(res, r - l)
			for l <= i {
				delete(m, s[l])
				l++
			}
		}

		m[s[r]] = r
		r++
	}

	return max(res, r - l)
}
