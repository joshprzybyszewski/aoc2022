package thirteen

import (
	"strconv"
	"strings"
)

func Two(
	input string,
) (string, error) {
	lines := strings.Split(input, "\n")

	return strconv.Itoa(len(lines)), nil
}
