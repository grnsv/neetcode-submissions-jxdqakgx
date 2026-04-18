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
		} else if nums[l] <= nums[m] && nums[l] < target && target < nums[m] ||
			nums[l] >= nums[m] && nums[l] > target && target < nums[m] {
			r = m
		} else if nums[m] <= nums[r] && nums[m] < target && target < nums[r] ||
			nums[m] >= nums[r] && nums[m] > target && target < nums[r] {
			l = m
		}

		l++
		r--
	}

	return -1
}
