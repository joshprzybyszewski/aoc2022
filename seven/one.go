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
) (int, error) {
	_, ds, err := getDirectorySizes(input)
	if err != nil {
		return 0, err
	}

	total := 0
	for _, size := range ds {
		if size <= part1max {
			total += size
		}
	}

	return total, nil
}

func getDirectorySizes(
	input string,
) (int, []int, error) {
	data := make([]fileDir, 1, 457)
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
	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		if nli == 0 {
			input = input[1:]
			continue
		}
		if input[0] == '$' {
			if isLS {
				data[curDirIndex].lsEndIndex = len(data)
			}
			isLS = false
			if input[2] == 'c' { // line starts with: "$ cd "
				curDirIndex = getNewIndexFromCD(input[:nli])
				if curDirIndex == -1 {
					return 0, nil, fmt.Errorf("invalid cd command: %q", input[:nli])
				}
			} else if input[2] == 'l' { // line starts with: "$ ls"
				if data[curDirIndex].lsIndex >= 0 {
					return 0, nil, fmt.Errorf("attempting to ls another time: %q", data[curDirIndex].name)
				}
				isLS = true
				data[curDirIndex].lsIndex = len(data)
			}
			input = input[nli+1:]
			continue
		}
		if !isLS {
			input = input[nli+1:]
			continue
		}

		if nli >= 4 && input[:4] == `dir ` {
			// this assumes that there cannot be a file named "dir".
			data = append(data, fileDir{
				name:       input[4:nli], // line starts with "dir "
				parent:     curDirIndex,
				size:       -1, // unset
				lsIndex:    -1, // unset
				lsEndIndex: -1, // unset
			})
			input = input[nli+1:]
			continue
		}

		size, err = strconv.Atoi(
			input[:strings.Index(input, ` `)],
		)
		if err != nil {
			return 0, nil, err
		}

		data = append(data, fileDir{
			size: size, // file size
			// name:   parts[1], // filename is unnecessary
			parent: -1, // files don't need to know the parent
		})
		input = input[nli+1:]
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

	rootSize, _ := getSize(rootDirIndex)
	dirSizes := make([]int, 0, 187)
	var ok bool
	for i := range data {
		size, ok = getSize(i)
		if ok {
			dirSizes = append(dirSizes, size)
		}
	}

	return rootSize, dirSizes, nil

}
