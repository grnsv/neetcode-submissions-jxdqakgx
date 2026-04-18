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
	s, ok := this.data[key]
	if !ok {
		return ""
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
			l = m + 1
		}
	}

	if l > 0 {
		return s[l-1].value
	}

	return ""
}
