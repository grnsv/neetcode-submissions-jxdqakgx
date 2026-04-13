func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	if n < 2 {
		return res
	}

	for i := n - 2; i >= 0; i-- {
		j := i + 1
		for j < n && temperatures[j] <= temperatures[i] {
			if res[j] == 0 {
				j = n
				break
			}

			j+=res[j]
		}

		if j < n {
			res[i] = j - i
		}
	}

	return res
}
