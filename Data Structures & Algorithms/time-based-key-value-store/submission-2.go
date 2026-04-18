type TimeMap struct {
	data map[string][]pair
}

type pair struct {
	timestamp int
	value     string
}

func Constructor() TimeMap {
	return TimeMap{data: make(map[string][]pair)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	s := this.data[key]
	s = append(s, pair{timestamp: timestamp, value: value})
	this.data[key] = s
}

func (this *TimeMap) Get(key string, timestamp int) string {
	res := ""
	s, ok := this.data[key]
	if !ok {
		return res
	}

	l := 0
	r := len(s) - 1
	for l <= r {
		m := l + (r-l)/2
		if s[m].timestamp == timestamp {
			return s[m].value
		}

		if timestamp < s[m].timestamp {
			r = m - 1
		} else {
			res = s[m].value
			l = m + 1
		}
	}

	return res
}
