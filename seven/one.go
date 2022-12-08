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
	err := db.processLines(lines)
	if err != nil {
		return ``, err
	}

	dirs := db.allDirsWithMaxSize(100000)

	total := 0
	for _, d := range dirs {
		total += d.size()
	}

	return strconv.Itoa(total), nil
}
