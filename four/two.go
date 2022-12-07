package four

import "fmt"

func Two(
	input string,
) (string, error) {
	ass, err := convertInputToAssignments(input)
	if err != nil {
		return ``, err
	}

	total := 0
	for _, as := range ass {
		if overlapping(as[0], as[1]) {
			total++
		}
	}

	return fmt.Sprintf("%d", total), nil
}
