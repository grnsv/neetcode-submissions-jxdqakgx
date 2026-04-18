func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		if nums[l] == target {
			return l
		} else if nums[r] == target {
			return r
		}

		m := l + (r-l)/2
		if nums[m] == target {
			return m
		}

		if target < nums[m] &&
			(nums[l] < target && nums[l] < nums[m] ||
				nums[l] > target && nums[l] > nums[m]) {
			r = m
		} else if target < nums[r] &&
			(nums[m] < target && nums[m] < nums[r] ||
				nums[m] > target && nums[m] > nums[r]) {
			l = m
		}

		l++
		r--
	}

	return -1
}
