import (
	"cmp"
	"slices"
)

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

	slices.SortFunc(pairs, func(a, b pair) int {
		return cmp.Compare(a.p, b.p)
	})

	fmt.Println(pairs)

	prev := float32(target - pairs[n - 1].p) / float32(pairs[n - 1].s)
	for i := n - 2; i >= 0; i-- {
		curr := max(prev, float32(target - pairs[i].p) / float32(pairs[i].s))
		if curr > prev {
			res++
			prev = curr
		}
	}

	return res
}
