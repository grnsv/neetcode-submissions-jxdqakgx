func lengthOfLongestSubstring(s string) int {
	res := 0
	for set, l, r := make(map[byte]struct{}), 0, 0; r < len(s); r++ {
		for _, exists := set[s[r]]; exists; _, exists = set[s[r]] {
			delete(set, s[l])
			l++
		}

		set[s[r]] = struct{}{}
		res = max(res, r-l+1)
	}

	return res
}
