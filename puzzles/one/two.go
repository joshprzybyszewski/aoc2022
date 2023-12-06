package one

import (
	"strings"
)

func Two(
	input string,
) (int, error) {
	sum := 0
	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		sum += getValueWithString(input[:nli])
		input = input[nli+1:]
	}

	return sum, nil
}

var stringValues = []string{
	`zero`,
	`one`,
	`two`,
	`three`,
	`four`,
	`five`,
	`six`,
	`seven`,
	`eight`,
	`nine`,
}

func getStringValue(input string) int {
	switch input[0] {
	case 'z', 'o', 't', 'f', 's', 'e', 'n':
		// these are the only valid starting chars
	default:
		return -1
	}

	for i, sv := range stringValues {
		if input == sv || (len(input) > len(sv) && input[:len(sv)] == sv) {
			return i
		}
	}
	return -1
}

func getValueWithString(line string) int {
	first := -1
	firstIndex := -1
	last := -1
	var i, val int
	var c byte

	for i = 0; i < len(line); i++ {
		c = line[i]
		if c >= '0' && c <= '9' {
			firstIndex = i
			first = int(c - '0')

			break
		}
		next := line[i:]
		if i+5 < len(line) {
			next = line[i : i+5]
		}

		val = getStringValue(next)
		if val == -1 {
			continue
		}

		firstIndex = i
		first = val
		break
	}
	if first == -1 {
		return 0
	}

	for i = len(line) - 1; i >= firstIndex; i-- {
		c = line[i]
		if c >= '0' && c <= '9' {

			last = int(c - '0')
			break
		}
		next := line[i:]
		if i+5 < len(line) {
			next = line[i : i+5]
		}

		val = getStringValue(next)
		if val == -1 {
			continue
		}
		last = val
		break
	}

	// last := -1

	// for i = len(line) - 1; i > first; i-- {
	// 	c = line[i]
	// 	if c >= '0' && c <= '9' {
	// 		last = int(c - '0')
	// 		break
	// 	}
	// 	next := line[i:]
	// 	if i+5 < len(line) {
	// 		next = line[i : i+5]
	// 	}

	// 	tmp = getStringValue(next)
	// 	if tmp == -1 {
	// 		continue
	// 	}
	// 	last = tmp
	// 	break
	// }
	// if last == -1 {
	// 	last = first
	// }

	return first*10 + last
}
