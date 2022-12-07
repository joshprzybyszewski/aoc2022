package seven

import (
	"fmt"
	"strings"
)

func Two(
	input string,
) (string, error) {
	lines := strings.Split(input, "\n")

	db := newDirBuilder()
	db.processLines(lines)

	dirs := db.allDirs()

	totalSpace := 70000000
	requiredUnusedSpace := 30000000
	curUsedSpace := db.root.size()

	dSize := -1

	for _, d := range dirs {
		mySize := d.size()
		newFreeSpace := totalSpace - (curUsedSpace - mySize)
		if newFreeSpace < requiredUnusedSpace {
			continue
		}
		fmt.Printf("found %q\n\twith size: %d\n", d.name, mySize)
		fmt.Printf("\tdeleting would bring free space to %d\n", newFreeSpace)

		if dSize == -1 || mySize < dSize {
			dSize = mySize
		}
	}
	if dSize == -1 {
		return ``, fmt.Errorf("answer not found")
	}

	return fmt.Sprintf("%d", dSize), nil
}
