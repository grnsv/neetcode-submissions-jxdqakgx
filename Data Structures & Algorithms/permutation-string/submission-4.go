func checkInclusion(s1 string, s2 string) bool {
	var letters [26]uint16
	for i := range len(s1) {
		letters[s1[i]-'a']++
	}

	var windowLetters [26]uint16
	l := 0
	for r := range len(s2) {
		i := s2[r]-'a'
		windowLetters[i]++
		if windowLetters[i] == letters[i] && r-l+1 == len(s1) {
			return true
		}

		if windowLetters[i] > letters[i] {
			for s2[l] != s2[r] {
				windowLetters[s2[l]-'a']--
				l++
			}

			windowLetters[s2[l]-'a']--
			l++
		}
	}

	return false
}
