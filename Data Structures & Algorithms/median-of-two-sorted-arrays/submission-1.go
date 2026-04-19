func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	a, b := nums1, nums2
	if len(b) < len(a) {
		a, b = b, a
	}

	total := len(a) + len(b)
	half := (total + 1) / 2

	l := -1
	r := len(a)
	var aLeft, aRight, bLeft, bRight int
	for {
		i := l + (r-l)/2
		j := half - i - 2

		if i >= 0 {
			aLeft = a[i]
		} else {
			aLeft = math.MinInt32
		}

		if i+1 < len(a) {
			aRight = a[i+1]
		} else {
			aRight = math.MaxInt32
		}

		if j >= 0 {
			bLeft = b[j]
		} else {
			bLeft = math.MinInt32
		}

		if j+1 < len(b) {
			bRight = b[j+1]
		} else {
			bRight = math.MaxInt32
		}

		if aLeft > bRight {
			r = i - 1
		} else if bLeft > aRight {
			l = i + 1
		} else {
			break
		}
	}

	if total%2 == 0 {
		return float64(max(aLeft, bLeft)+min(aRight, bRight)) / 2.0
	}

	return float64(max(aLeft, bLeft))
}
