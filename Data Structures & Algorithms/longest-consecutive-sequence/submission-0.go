func longestConsecutive(nums []int) int {
	set := make(map[int]bool, len(nums)/2)
	for _, num := range nums {
		set[num] = true
	}

	maxLen := 0
	for num, mustSee := range set {
		if !mustSee {
			continue
		}

		length := 1
		for curr := num; ; {
			curr--
			mustSee = set[curr]
			if !mustSee {
				break
			}

			set[curr] = false
			length++
		}

		for curr := num; ; {
			curr++
			mustSee = set[curr]
			if !mustSee {
				break
			}

			set[curr] = false
			length++
		}

		maxLen = max(maxLen, length)
	}

	return maxLen
}
