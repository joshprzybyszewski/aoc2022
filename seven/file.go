package seven

import (
	"fmt"
	"strconv"
	"strings"
)

type file struct {
	name string
	size int
}

func newFile(line string) (*file, error) {
	parts := strings.Split(line, ` `)
	if len(parts) != 2 {
		return nil, fmt.Errorf("line should have two parts: %q", line)
	}

	size, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}

	return &file{
		name: parts[1],
		size: size,
	}, nil
}
