func minEatingSpeed(piles []int, h int) int {
	var l, r int

	sum := 0
	for _, pile := range piles {
		sum += pile
		r = max(r, pile)
	}

	l = sum / h
	if sum % h > 0 {
		l++
	}

	k := r
	for l <= r {
		m := (l + r) / 2
		t := 0
		for _, p := range piles {
			t += p / m
			if p % m > 0 {
				t++
			}
		}
		if t <= h {
			r = m - 1
			k = min(k, m)
		} else {
			l = m + 1
		}
	}

	return k
}
