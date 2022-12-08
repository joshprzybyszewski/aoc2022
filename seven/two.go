package seven

import (
	"fmt"
	"strconv"
	"strings"
)

func Two(
	input string,
) (string, error) {
	lines := strings.Split(input, "\n")

	db := newDirBuilder()
	err := db.processLines(lines)
	if err != nil {
		return ``, err
	}

	totalSpace := 70000000
	requiredUnusedSpace := 30000000
	curUsedSpace := db.root.size()

	dSize := db.root.getMinSizeGreaterThan(
		requiredUnusedSpace - (totalSpace - curUsedSpace),
	)
	if dSize == -1 {
		return ``, fmt.Errorf("answer not found")
	}

	return strconv.Itoa(dSize), nil
}
