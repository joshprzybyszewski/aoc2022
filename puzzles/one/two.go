package one

import (
	"bytes"
	"strings"
)

func Two(
	input string,
) (int, error) {
	sum := 0
	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		sum += getValueWithString([]byte(input[:nli]))
		input = input[nli+1:]
	}

	return sum, nil
}

var stringValues = [][]byte{
	[]byte(`zero`),
	[]byte(`one`),
	[]byte(`two`),
	[]byte(`three`),
	[]byte(`four`),
	[]byte(`five`),
	[]byte(`six`),
	[]byte(`seven`),
	[]byte(`eight`),
	[]byte(`nine`),
}

func getStringValue(val []byte) int {
	for i, sv := range stringValues {
		val := val
		if len(val) > len(sv) {
			val = val[:len(sv)]
		}
		if bytes.Equal(val, sv) {
			return i

		}
	}
	return -1
}

func getValueWithString(line []byte) int {
	first, last := -1, -1

	for i, c := range line {
		if c >= '0' && c <= '9' {
			if first == -1 {
				first = int(c - '0')
			}
			last = int(c - '0')
			continue
		}
		next := line[i:]
		if i+5 < len(line) {
			next = line[i : i+5]
		}

		val := getStringValue(next)
		if val == -1 {
			continue
		}
		if first == -1 {
			first = val
		}
		last = val

	}
	if first == -1 {
		return 0
	}

	return first*10 + last
}
