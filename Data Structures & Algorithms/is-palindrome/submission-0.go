func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	for l < r {
		if (s[l] < 'A' || s[l] > 'Z') && (s[l] < 'a' || s[l] > 'z') && (s[l] < '0' || s[l] > '9') {
			l++
			continue
		}

		if (s[r] < 'A' || s[r] > 'Z') && (s[r] < 'a' || s[r] > 'z') && (s[r] < '0' || s[r] > '9') {
			r--
			continue
		}

		if strings.ToLower(string(s[l])) != strings.ToLower(string(s[r])) {
			return false
		}

		l++
		r--
	}

	return true
}
