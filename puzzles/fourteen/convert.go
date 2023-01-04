package fourteen

import (
	"strconv"
	"strings"
)

func getGrid(input string) (grid, error) {
	g := newGrid()

	var coords []coord
	var err error
	var i int
	var c, next coord

	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		if nli == 0 {
			input = input[nli+1:]
			continue
		}

		coords, err = getCoords(input[:nli])
		if err != nil {
			return grid{}, err
		}
		for i = 0; i < len(coords)-1; i++ {
			c = coords[i]
			next = coords[i+1]
			for {
				g.addRock(c.x, c.y)

				if c.x < next.x {
					c.x++
				} else if c.x > next.x {
					c.x--
				} else if c.y < next.y {
					c.y++
				} else if c.y > next.y {
					c.y--
				} else {
					break
				}
			}
		}
		input = input[nli+1:]
	}

	return g, nil
}

func getCoords(line string) ([]coord, error) {
	const del = " -> "
	var ci, x, y int
	var err error

	output := make([]coord, 0, 4)
	for deli := strings.Index(line, del); len(line) > 0; deli = strings.Index(line, del) {
		if deli == 0 {
			line = line[deli+len(del):]
			continue
		}

		ci = strings.Index(line, `,`)
		x, err = strconv.Atoi(line[:ci])
		if err != nil {
			return nil, err
		}
		if deli == -1 {
			y, err = strconv.Atoi(line[ci+1:])
			line = line[:0]
		} else {
			y, err = strconv.Atoi(line[ci+1 : deli])
			line = line[deli+len(del):]
		}
		if err != nil {
			return nil, err
		}

		output = append(output, coord{
			x: x,
			y: y,
		})
	}

	return output, nil
}
