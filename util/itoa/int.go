package itoa

func Int(s string) int {
	if s[0] == '-' {
		// the first byte was a dash character and signals a negative number
		out := int(s[1] - '0')
		for i := 2; i < len(s); i++ {
			out *= 10
			out += int(s[i] - '0')
		}
		return -out
	}

	out := int(s[0] - '0')
	for i := 1; i < len(s); i++ {
		out *= 10
		out += int(s[i] - '0')
	}

	return out
}