package four

import (
	"fmt"
	"strings"
)

func One(
	input string,
) (string, error) {
	lines := strings.Split(input, "\n")

	pairs, err := convertInputToAssignments(lines)
	if err != nil {
		return ``, err
	}

	total := 0
	for _, p := range pairs {
		if p.fullyContained() {
			total++
		}
	}

	return fmt.Sprintf("%d", total), nil
}

func convertInputToAssignments(
	lines []string,
) ([]pair, error) {

	output := make([]pair, 0, len(lines))

	for _, line := range lines {
		if line == `` {
			continue
		}
		p, err := newPair(line)
		if err != nil {
			return nil, err
		}
		output = append(output, p)
	}

	return output, nil
}
