package seven

import (
	"strconv"
	"strings"
)

const (
	rootDirName = `/`
)

func One(
	input string,
) (string, error) {
	lines := strings.Split(input, "\n")

	db := newDirBuilder()
	db.processLines(lines)

	dirs := db.allDirsWithMaxSize(100000)

	total := 0
	for _, d := range dirs {
		total += d.size()
	}

	return strconv.Itoa(total), nil
}
