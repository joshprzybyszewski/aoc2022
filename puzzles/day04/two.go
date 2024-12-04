package day04

const (
	X_S_ul_SET uint8 = 1 << 0
	X_S_ur_SET uint8 = 1 << 1
	X_S_dl_SET uint8 = 1 << 2
	X_S_dr_SET uint8 = 1 << 3
	X_M_ul_SET uint8 = 1 << 4
	X_M_ur_SET uint8 = 1 << 5
	X_M_dl_SET uint8 = 1 << 6
	X_M_dr_SET uint8 = 1 << 7

	XMAS_1 uint8 = X_S_ul_SET | X_S_ur_SET | X_M_dl_SET | X_M_dr_SET
	XMAS_2 uint8 = X_S_ur_SET | X_S_dr_SET | X_M_dl_SET | X_M_ul_SET
	XMAS_3 uint8 = X_S_dr_SET | X_S_dl_SET | X_M_ul_SET | X_M_ur_SET
	XMAS_4 uint8 = X_S_dl_SET | X_S_ul_SET | X_M_ur_SET | X_M_dr_SET
)

type xmasCoord struct {
	isA    bool
	nearby uint8
}

func Two(
	inputString string,
) (int, error) {

	input := []byte(inputString)
	grid := [GRID_SIZE][GRID_SIZE]xmasCoord{}

	var r, c, i int
	setM := func() {
		if c > 0 {
			if r > 0 {
				grid[r-1][c-1].nearby |= X_M_dr_SET
			}
			if r < len(grid)-1 {
				grid[r+1][c-1].nearby |= X_M_ur_SET
			}
		}
		if c < len(grid[r])-1 {
			if r > 0 {
				grid[r-1][c+1].nearby |= X_M_dl_SET
			}
			if r < len(grid)-1 {
				grid[r+1][c+1].nearby |= X_M_ul_SET
			}
		}
	}
	setA := func() {
		grid[r][c].isA = true
	}
	setS := func() {
		if c > 0 {
			if r > 0 {
				grid[r-1][c-1].nearby |= X_S_dr_SET
			}
			if r < len(grid)-1 {
				grid[r+1][c-1].nearby |= X_S_ur_SET
			}
		}
		if c < len(grid[r])-1 {
			if r > 0 {
				grid[r-1][c+1].nearby |= X_S_dl_SET
			}
			if r < len(grid)-1 {
				grid[r+1][c+1].nearby |= X_S_ul_SET
			}
		}
	}

	var b byte
	for r = 0; r < len(grid); r++ {
		for c = 0; c < len(grid[r]); c++ {
			b = input[i]
			i++
			switch b {
			case 'X':
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
			if !grid[r][c].isA {
				continue
			}
			if grid[r][c].nearby == XMAS_1 ||
				grid[r][c].nearby == XMAS_2 ||
				grid[r][c].nearby == XMAS_3 ||
				grid[r][c].nearby == XMAS_4 {
				total++
			}
		}
	}

	// 36 is not the right answer

	return total, nil
}
