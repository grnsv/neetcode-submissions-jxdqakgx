import "slices"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		bytes := []byte(str)
		slices.Sort(bytes)
		key := string(bytes)
		strs = m[key]
		m[key] = append(strs, str)
	}
	res := make([][]string, 0, len(m))
	for _, strs := range m {
		res = append(res, strs)
	}

	return res
}
