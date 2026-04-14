func carFleet(target int, position []int, speed []int) int {
	res := 1
	n := len(position)
	if n == 1 {
		return res
	}

	type pair struct {
		p int
		s int
	}
	pairs := make([]pair, n)
	for i := range n {
		pairs[i] = pair{p: position[i], s: speed[i]}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].p > pairs[j].p
	})

	prev := float32(target - pairs[0].p) / float32(pairs[0].s)
	for i := 1; i < n; i++ {
		curr := float32(target - pairs[i].p) / float32(pairs[i].s)
		if curr > prev {
			res++
			prev = curr
		}
	}

	return res
}
