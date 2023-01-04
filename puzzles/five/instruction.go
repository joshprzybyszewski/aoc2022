package five

import (
	"strconv"
	"strings"
)

type instruction struct {
	source   int
	dest     int
	quantity int
}

func newInstruction(line string) (instruction, error) {
	// The `line` looks like:
	// "move 6 from 6 to 5"
	// Using Sscanf allocates a lot of bytes to the heap, so
	// we avoid that with this indexing weirdness.
	var q, s, d int
	var err error

	i1 := 5
	i2 := i1 + strings.Index(line[i1:], ` `)
	q, err = strconv.Atoi(line[i1:i2])
	if err != nil {
		return instruction{}, err
	}

	i1 = i2 + 6 // skip past " from "
	i2 = i1 + strings.Index(line[i1:], ` `)

	s, err = strconv.Atoi(line[i1:i2])
	if err != nil {
		return instruction{}, err
	}

	i1 = i2 + 4 // skip past " to "
	// the last number goes to the end of the line, so we don't use i2 again

	d, err = strconv.Atoi(line[i1:])
	if err != nil {
		return instruction{}, err
	}

	return instruction{
		quantity: q,
		source:   s,
		dest:     d,
	}, nil
}
