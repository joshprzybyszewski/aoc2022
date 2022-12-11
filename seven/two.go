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

	data := make([]fileDir, 1, 1+len(lines))
	const rootDirIndex = 0
	const rootName = `/`
	data[rootDirIndex] = fileDir{
		lsIndex:    -1, // unset
		lsEndIndex: -1, // unset
		size:       -1, // unset
		name:       rootName,
		parent:     -1,
	}

	curDirIndex := rootDirIndex

	var tmp int
	getNewIndexFromCD := func(line string) int {
		tmp = strings.LastIndex(line, ` `) + 1
		if line[tmp:] == rootName {
			return rootDirIndex
		}
		if line[tmp:] == `..` {
			return data[curDirIndex].parent
		}

		for i := data[curDirIndex].lsIndex; i < data[curDirIndex].lsEndIndex; i++ {
			if data[i].parent != -1 && line[tmp:] == data[i].name {
				// it's a directory with the desired name
				return i
			}
		}

		return -1
	}

	isLS := false
	var size int
	var err error
	for _, line := range lines {
		if line == `` {
			continue
		}
		if line[0] == '$' {
			if isLS {
				data[curDirIndex].lsEndIndex = len(data)
			}
			isLS = false
			if line[2] == 'c' { // line starts with: "$ cd "
				curDirIndex = getNewIndexFromCD(line)
				if curDirIndex == -1 {
					return ``, fmt.Errorf("invalid cd command: %q", line)
				}
			} else if line[2] == 'l' { // line starts with: "$ ls"
				if data[curDirIndex].lsIndex >= 0 {
					return ``, fmt.Errorf("attempting to ls another time: %q", data[curDirIndex].name)
				}
				isLS = true
				data[curDirIndex].lsIndex = len(data)
			}
			continue
		}
		if !isLS {
			continue
		}
		if isDir(line) {
			data = append(data, fileDir{
				name:       line[4:], // line starts with "dir "
				parent:     curDirIndex,
				size:       -1, // unset
				lsIndex:    -1, // unset
				lsEndIndex: -1, // unset
			})
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
			parent: -1, // files don't need to know the parent
		})
	}

	if isLS {
		data[curDirIndex].lsEndIndex = len(data)
		isLS = false
	}

	var getSize func(fdi int) (int, bool)
	getSize = func(fdi int) (int, bool) {
		if data[fdi].parent < 0 && fdi != rootDirIndex {
			// not a child directory, and not the root => it's just a file
			return data[fdi].size, false
		}
		if data[fdi].size >= 0 {
			// size has been set already
			return data[fdi].size, true
		}

		// `total` and `innerSize` need to be scoped inside this function because it's recursive
		total := 0
		var innerSize int
		for i := data[fdi].lsIndex; i < data[fdi].lsEndIndex; i++ {
			innerSize, _ = getSize(i)
			total += innerSize
		}
		data[fdi].size = total
		return total, true
	}

	totalSpace := 70000000
	requiredUnusedSpace := 30000000
	curUsedSpace, _ := getSize(rootDirIndex)

	dSize := totalSpace + 1
	min := requiredUnusedSpace - (totalSpace - curUsedSpace)
	var ok bool
	for i := range data {
		size, ok = getSize(i)
		if !ok || size < min {
			// not a directory OR
			// not large enough
			continue
		}
		if size < dSize {
			dSize = size
		}
	}

	if dSize > totalSpace {
		return ``, fmt.Errorf("answer not found")
	}

	return strconv.Itoa(dSize), nil
}
