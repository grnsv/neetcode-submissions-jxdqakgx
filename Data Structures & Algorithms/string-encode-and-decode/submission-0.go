type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	sb := strings.Builder{}
	l := 0
	for _, str := range strs {
		l += len(str) + 2
	}
	sb.Grow(l)
	for _, str := range strs {
		sb.WriteString(strconv.Itoa(len(str)))
		sb.WriteString("#")
		sb.WriteString(str)
	}

	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	var strs []string
	i := 0
	for i < len(encoded) {
		l := 0
		for encoded[i] != '#' {
			l *= 10
			l += int(encoded[i] - '0')
			i++
		}
		i++
		str := string(encoded[i:i+l])
		strs = append(strs, str)
		i += l
	}
	return strs
}
