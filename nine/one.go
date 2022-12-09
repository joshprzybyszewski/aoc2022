package nine

import (
	"strconv"
	"strings"
)

func One(
	input string,
) (string, error) {
	lines := strings.Split(input, "\n")

	r := rope{}
	var q, i int
	var err error

	pos := make(map[coord]struct{}, len(lines))
	pos[r.tail] = struct{}{}

	for _, line := range lines {
		if line == `` {
			continue
		}
		q, err = strconv.Atoi(line[2:])
		if err != nil {
			return ``, err
		}

		for i = 0; i < q; i++ {
			r = move(r, direction(line[0]))
			pos[r.tail] = struct{}{}
		}
	}

	return strconv.Itoa(len(pos)), nil
}

type coord struct {
	x, y int
}

type rope struct {
	head coord
	tail coord
}

type direction byte

const (
	right direction = 'R'
	up    direction = 'U'
	left  direction = 'L'
	down  direction = 'D'
)

func move(
	r rope,
	dir direction,
) rope {
	switch dir {
	case right:
		r.head.x++
	case left:
		r.head.x--
	case up:
		r.head.y++
	case down:
		r.head.y--
	}

	r.tail = moveCoord(r.tail, r.head)

	return r
}
