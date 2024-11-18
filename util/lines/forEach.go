package lines

import "strings"

func ForEach(
	input string,
	every func(line string),
) {
	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		every(input[:nli])
		input = input[nli+1:]
	}
}
