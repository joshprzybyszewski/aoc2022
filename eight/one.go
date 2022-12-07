package eight

import (
	"fmt"
	"strings"
)

const (
	rootDirName = `/`
)

func One(
	input string,
) (string, error) {
	lines := strings.Split(input, "\n")

	return fmt.Sprintf("%d", len(lines)), nil
}
