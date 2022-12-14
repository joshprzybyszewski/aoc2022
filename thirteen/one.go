package thirteen

import (
	"strconv"
	"strings"
)

func One(
	input string,
) (string, error) {
	lines := strings.Split(input, "\n")

	return strconv.Itoa(len(lines)), nil
}
