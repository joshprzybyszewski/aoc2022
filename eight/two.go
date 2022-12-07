package eight

import (
	"fmt"
	"strings"
)

func Two(
	input string,
) (string, error) {
	lines := strings.Split(input, "\n")

	return fmt.Sprintf("%d", len(lines)), nil
}
