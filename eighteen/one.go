package eighteen

import (
	"strconv"
	"strings"
)

const (
	dropletSideLength = 22
)

type droplet [dropletSideLength][dropletSideLength][dropletSideLength]uint8

func (d *droplet) fill(
	x, y, z int,
) {
	d[x][y][z] = 6

	if z > 0 && d[x][y][z-1] > 0 {
		d[x][y][z-1]--
		d[x][y][z]--
	}

	if z < dropletSideLength-1 && d[x][y][z+1] > 0 {
		d[x][y][z+1]--
		d[x][y][z]--
	}

	if y > 0 && d[x][y-1][z] > 0 {
		d[x][y-1][z]--
		d[x][y][z]--
	}

	if y < dropletSideLength-1 && d[x][y+1][z] > 0 {
		d[x][y+1][z]--
		d[x][y][z]--
	}

	if x > 0 && d[x-1][y][z] > 0 {
		d[x-1][y][z]--
		d[x][y][z]--
	}

	if x < dropletSideLength-1 && d[x+1][y][z] > 0 {
		d[x+1][y][z]--
		d[x][y][z]--
	}

}

func One(
	input string,
) (int, error) {
	d := droplet{}

	var i1, i2 int
	var x, y, z int
	var err error

	for nli := strings.Index(input, "\n"); nli >= 0; nli = strings.Index(input, "\n") {
		if nli == 0 {
			input = input[nli+1:]
			continue
		}

		i1 = 0
		i2 = strings.Index(input, `,`)
		x, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return 0, err
		}

		i1 = i2 + 1
		i2 = i1 + strings.Index(input[i1:], `,`)
		y, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return 0, err
		}

		i1 = i2 + 1
		i2 = nli
		z, err = strconv.Atoi(input[i1:i2])
		if err != nil {
			return 0, err
		}
		d.fill(x, y, z)

		input = input[nli+1:]
	}

	total := 0
	for xi := range d {
		for yi := range d[xi] {
			for zi := range d[xi][yi] {
				total += int(d[xi][yi][zi])
			}
		}
	}

	return total, nil
}
