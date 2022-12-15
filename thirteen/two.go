package thirteen

import (
	"strings"
)

func Two(
	input string,
) (int, error) {
	lines := strings.Split(input, "\n")

	return len(lines), nil
}
