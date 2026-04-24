func maxSlidingWindow(nums []int, k int) []int {
    res := make([]int, 0, len(nums) - k + 1)
	q := []int{}
	for i, num := range nums {
		for len(q) > 0 && nums[q[len(q)-1]] <= num {
			q = q[:len(q)-1]
		}

		q = append(q, i)

		if q[0] == i - k {
			q = q[1:]
		}

		if i >= k - 1 {
			res = append(res, nums[q[0]])
		}
	}

	return res
}
