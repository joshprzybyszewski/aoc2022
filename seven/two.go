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

	curUsedSpace, ds, err := getDirectorySizes(lines)
	if err != nil {
		return ``, err
	}

	totalSpace := 70000000
	requiredUnusedSpace := 30000000
	min := requiredUnusedSpace - (totalSpace - curUsedSpace)

	smallestDirSize := totalSpace + 1
	for _, size := range ds {
		if size >= min && size < smallestDirSize {
			smallestDirSize = size
		}
	}

	if smallestDirSize > totalSpace {
		return ``, fmt.Errorf("answer not found")
	}

	return strconv.Itoa(smallestDirSize), nil
}
