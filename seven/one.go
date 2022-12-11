package seven

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	part1max    = 100000
	rootDirName = `/`
)

type fileDir struct {
	size int
	// if a directory, this is the index of the first file
	lsIndex, lsEndIndex int
	name                string
	// if a file, this is -1
	// if a directory, this is the index of its parent directory
	parent int
}

func One(
	input string,
) (string, error) {
	lines := strings.Split(input, "\n")

	data := make([]fileDir, 1, 1+len(lines))
	rootDirIndex := 0
	rootName := `/`
	data[rootDirIndex] = fileDir{
		lsIndex:    -1, // unset
		lsEndIndex: -1, // unset
		size:       -1, // unset
		name:       rootName,
		parent:     -1,
	}

	curDirIndex := rootDirIndex

	getChildIndex := func(name string) int {
		for i := data[curDirIndex].lsIndex; i < data[curDirIndex].lsEndIndex; i++ {
			if name == data[i].name {
				return i
			}
		}

		return -1
	}

	var getSize func(fdi int) (int, bool)
	var size int
	getSize = func(fdi int) (int, bool) {
		if data[fdi].lsIndex < 0 {
			// not a directory
			return data[fdi].size, false
		}
		if data[fdi].size >= 0 {
			// it's been set already
			return data[fdi].size, true
		}

		// these need to be scoped inside this function because it's recursive
		total := 0
		for i := data[fdi].lsIndex; i < data[fdi].lsEndIndex; i++ {
			size, _ = getSize(i)
			total += size
		}
		data[fdi].size = total
		return total, true
	}

	finishLS := func() {
		if data[curDirIndex].lsIndex >= 0 && data[curDirIndex].lsEndIndex == -1 {
			data[curDirIndex].lsEndIndex = len(data)
		}
	}

	var newDir string
	handleCD := func(line string) error {
		finishLS()

		newDir = line[strings.LastIndex(line, ` `)+1:]
		if newDir == rootName {
			curDirIndex = rootDirIndex
			return nil
		}
		if newDir == `..` {
			curDirIndex = data[curDirIndex].parent
			return nil
		}

		curDirIndex = getChildIndex(newDir)
		if curDirIndex == -1 {
			return fmt.Errorf("does not have a child with name: %q", newDir)
		}
		return nil
	}

	handleLS := func() {
		finishLS()
		data[curDirIndex].lsIndex = len(data)
	}

	var err error
	for _, line := range lines {
		if line == `` {
			continue
		}
		if isDir(line) {
			data = append(data, fileDir{
				size:       -1, // unset
				lsIndex:    -1, // unset
				lsEndIndex: -1, // unset
				name:       line[4:],
				parent:     curDirIndex,
			})
			continue
		}
		if isCDCmd(line) {
			handleCD(line)
			continue
		}
		if isLSCmd(line) {
			handleLS()
			continue
		}

		size, err = strconv.Atoi(
			line[:strings.Index(line, ` `)],
		)
		if err != nil {
			return ``, err
		}

		data = append(data, fileDir{
			size: size, // file size
			// name:   parts[1], // filename is unnecessary
			parent:     curDirIndex,
			lsIndex:    -1, // not a directory
			lsEndIndex: -1, // not a directory
		})
	}

	total := 0
	var ok bool
	for i := range data {
		size, ok = getSize(i)
		if ok && size <= part1max {
			total += size
		}
	}

	return strconv.Itoa(total), nil
}
