func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if j, ok := m[diff]; ok {
			return []int{min(i, j), max(i, j)}
		}
		m[num] = i
	}
	return []int{-1, -1}
}
