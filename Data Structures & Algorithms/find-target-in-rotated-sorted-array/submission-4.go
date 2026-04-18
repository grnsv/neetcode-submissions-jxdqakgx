func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		}

		if nums[l] <= nums[m] {
			if nums[m] < target || target < nums[l] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			if target < nums[m] || nums[r] < target {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}

	return -1
}
