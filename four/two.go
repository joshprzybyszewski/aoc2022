package four

import (
	"fmt"
	"strings"
)

func Two(
	input string,
) (string, error) {
	lines := strings.Split(input, "\n")

	pairs, err := convertInputToAssignments(lines)
	if err != nil {
		return ``, err
	}

	total := 0
	for _, p := range pairs {
		if p.overlapping() {
			total++
		}
	}

	return fmt.Sprintf("%d", total), nil
}
