func topKFrequent(nums []int, k int) []int {
	numFreqs := make(map[int]int)
	maxFreq := 0
	for _, num := range nums {
		freq := numFreqs[num] +1
		numFreqs[num] = freq
		maxFreq = max(maxFreq, freq)
	}

	freqNums := make(map[int][]int)
	for num, freq := range numFreqs {
		nums := freqNums[freq]
		freqNums[freq] = append(nums, num)
	}

	res := make([]int, k)
	i := 0
	for {
		nums := freqNums[maxFreq]
		for _, num := range nums {
			res[i] = num
			i++
			if i == k {
				return res
			}
		}
		maxFreq--
	}

	return []int{}
}
