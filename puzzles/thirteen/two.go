package thirteen

import (
	"strings"
)

func Two(
	input string,
) (int, error) {
	lines := strings.Split(input, "\n")

	marker1 := []interface{}{
		[]interface{}{
			2,
		},
	}
	marker2 := []interface{}{
		[]interface{}{
			6,
		},
	}

	m1, m2 := 1, 1

	var checkM1First bool
	switch compare(marker1, marker2) {
	case valid:
		checkM1First = false
		m2++
	default:
		checkM1First = true
		m1++
	}

	var l interface{}
	for _, line := range lines {
		if line == `` {
			continue
		}

		l, _ = parse(line, 0)

		if checkM1First {
			switch compare(l, marker1) {
			case valid:
				m1++
				switch compare(l, marker2) {
				case valid:
					m2++
				}
			}
		} else {
			switch compare(l, marker2) {
			case valid:
				m2++

				switch compare(l, marker1) {
				case valid:
					m1++
				}
			}
		}
	}

	return m1 * m2, nil
}
