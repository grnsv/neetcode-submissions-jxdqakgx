type TimeMap struct {
	data map[string]map[int]string
}

func Constructor() TimeMap {
	return TimeMap{data: make(map[string]map[int]string)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	m, ok := this.data[key]
	if !ok {
		m = make(map[int]string)
	}

	m[timestamp] = value
	this.data[key] = m
}

func (this *TimeMap) Get(key string, timestamp int) string {
	m, ok := this.data[key]
	if !ok {
		return ""
	}

	for timestamp > 0 {
		if value, ok := m[timestamp]; ok {
			return value
		}
		
		timestamp--
	}

	return ""
}
