package day04

const (
	GRID_SIZE = 140
)

const (
	X_SET   uint8 = 1 << 0
	M_SET   uint8 = 1 << 1
	A_SET   uint8 = 1 << 2
	S_SET   uint8 = 1 << 3
	ALL_SET uint8 = X_SET | M_SET | A_SET | S_SET
)

const (
	right uint8 = 0
	ur    uint8 = 1
	up    uint8 = 2
	ul    uint8 = 3
	left  uint8 = 4
	dl    uint8 = 5
	down  uint8 = 6
	dr    uint8 = 7
)

type coord struct {
	dirs [8]uint8
}

func (c *coord) setX(i uint8) {
	c.dirs[i] |= X_SET
}

func (c *coord) setM(i uint8) {
	c.dirs[i] |= M_SET
}

func (c *coord) setA(i uint8) {
	c.dirs[i] |= A_SET
}

func (c *coord) setS(i uint8) {
	c.dirs[i] |= S_SET
}

func One(
	inputString string,
) (int, error) {

	input := []byte(inputString)
	grid := [GRID_SIZE][GRID_SIZE]coord{}

	var r, c, i int
	setX := func() {
		grid[r][c].setX(right)
		grid[r][c].setX(ur)
		grid[r][c].setX(up)
		grid[r][c].setX(ul)
		grid[r][c].setX(left)
		grid[r][c].setX(dl)
		grid[r][c].setX(down)
		grid[r][c].setX(dr)
	}
	setM := func() {
		if c > 0 {
			grid[r][c-1].setM(right)
			if r > 0 {
				grid[r-1][c-1].setM(dr)
			}
			if r < len(grid)-1 {
				grid[r+1][c-1].setM(ur)
			}
		}
		if c < len(grid[r])-1 {
			grid[r][c+1].setM(left)
			if r > 0 {
				grid[r-1][c+1].setM(dl)
			}
			if r < len(grid)-1 {
				grid[r+1][c+1].setM(ul)
			}
		}
		if r > 0 {
			grid[r-1][c].setM(down)
		}
		if r < len(grid)-1 {
			grid[r+1][c].setM(up)
		}
	}
	setA := func() {
		if c > 1 {
			grid[r][c-2].setA(right)
			if r > 1 {
				grid[r-2][c-2].setA(dr)
			}
			if r < len(grid)-2 {
				grid[r+2][c-2].setA(ur)
			}
		}
		if c < len(grid[r])-2 {
			grid[r][c+2].setA(left)
			if r > 1 {
				grid[r-2][c+2].setA(dl)
			}
			if r < len(grid)-2 {
				grid[r+2][c+2].setA(ul)
			}
		}
		if r > 1 {
			grid[r-2][c].setA(down)
		}
		if r < len(grid)-2 {
			grid[r+2][c].setA(up)
		}
	}
	setS := func() {
		if c > 2 {
			grid[r][c-3].setS(right)
			if r > 2 {
				grid[r-3][c-3].setS(dr)
			}
			if r < len(grid)-3 {
				grid[r+3][c-3].setS(ur)
			}
		}
		if c < len(grid[r])-3 {
			grid[r][c+3].setS(left)
			if r > 2 {
				grid[r-3][c+3].setS(dl)
			}
			if r < len(grid)-3 {
				grid[r+3][c+3].setS(ul)
			}
		}
		if r > 2 {
			grid[r-3][c].setS(down)
		}
		if r < len(grid)-3 {
			grid[r+3][c].setS(up)
		}
	}

	var b byte
	for r = 0; r < len(grid); r++ {
		for c = 0; c < len(grid[r]); c++ {
			b = input[i]
			i++
			switch b {
			case 'X':
				setX()
			case 'M':
				setM()
			case 'A':
				setA()
			case 'S':
				setS()
				// default:
				// 	fmt.Printf("Found %q at (%d, %d)\n", string(b), r, c)
				// 	panic(`ahh` + string(b))
			}
		}
		// if input[i] != '\n' {
		// 	panic(`ahh`)
		// }
		// input[i] should be a newline
		i++
	}

	total := 0
	for r := 0; r < len(grid); r++ {
		for c = 0; c < len(grid[r]); c++ {
			for i = 0; i < len(grid[r][c].dirs); i++ {
				if grid[r][c].dirs[i] == ALL_SET {
					total++
				}
			}
		}
	}

	return total, nil
}
