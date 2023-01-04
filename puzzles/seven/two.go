package seven

import (
	"fmt"
)

func Two(
	input string,
) (int, error) {
	curUsedSpace, ds, err := getDirectorySizes(input)
	if err != nil {
		return 0, err
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
		return 0, fmt.Errorf("answer not found")
	}

	return smallestDirSize, nil
}
