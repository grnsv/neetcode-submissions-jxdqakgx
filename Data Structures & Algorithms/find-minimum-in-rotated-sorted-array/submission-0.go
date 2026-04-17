func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		if nums[l] < nums[r] {
			return nums[l]
		}
		m := l + (r - l) / 2
		if nums[l] > nums[m] {
			r = m
		} else if nums[m] > nums[r] {
			l = m + 1
		}
	}

	return nums[l]
}
