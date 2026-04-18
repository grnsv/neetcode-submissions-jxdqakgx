func productExceptSelf(nums []int) []int {
	prefixes := make([]int, len(nums) - 1)
	curr := 1
	for i := range len(nums) - 1 {
		num := nums[i]
		curr *= num
		prefixes[i] = curr
	}

	suffixes := make([]int, len(nums))
	curr = 1
	for i := len(nums) - 1; i > 0; i-- {
		num := nums[i]
		curr *= num
		suffixes[i] = curr
	}

	output := make([]int, len(nums))
	for i := range nums {
		prefix, suffix := 1, 1
		if i > 0 {
			prefix = prefixes[i-1]
		}
		if i < len(nums) - 1 {
			suffix = suffixes[i+1]
		}

		output[i] = prefix * suffix
	}

	return output
}
