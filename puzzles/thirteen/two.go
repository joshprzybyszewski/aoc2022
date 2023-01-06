package thirteen

import "strings"

func Two(
	input string,
) (int, error) {
	marker1 := []thing{{
		slice: []thing{{
			value: 2,
		}},
	}}
	marker2 := []thing{{
		slice: []thing{{
			value: 6,
		}},
	}}

	m1, m2 := 1, 1

	var checkM1First bool
	switch compareSlices(marker1, marker2) {
	case valid:
		checkM1First = false
		m2++
	default:
		checkM1First = true
		m1++
	}

	var l []thing
	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		if nli == 0 {
			// skip empty lines
			input = input[1:]
			continue
		}

		l, _ = parse(input[:nli], 0)
		input = input[nli+1:]

		if checkM1First {
			switch compareSlices(l, marker1) {
			case valid:
				m1++
				switch compareSlices(l, marker2) {
				case valid:
					m2++
				}
			}
		} else {
			switch compareSlices(l, marker2) {
			case valid:
				m2++

				switch compareSlices(l, marker1) {
				case valid:
					m1++
				}
			}
		}
	}

	return m1 * m2, nil
}
