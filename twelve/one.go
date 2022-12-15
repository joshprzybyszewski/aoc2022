package twelve

const (
	numRows = uint8(41)
	numCols = uint8(161)
)

type coord struct {
	row, col uint8
}

// [row][col]
type grid [numRows][numCols]uint8

func newGrid(input string) (grid, coord, coord) {
	var s, e coord
	var g grid
	var i int
	var b byte
	for r := uint8(0); r < numRows; r++ {
		for c := uint8(0); c < numCols; c++ {
			b = input[i]
			if b == 'E' {
				e = coord{
					row: r,
					col: c,
				}
				g[r][c] = 25
			} else if b == 'S' {
				s = coord{
					row: r,
					col: c,
				}
				g[r][c] = 0
			} else {
				g[r][c] = b - 'a'
			}
			i++
		}
		// skip the newline char
		i++
	}
	return g, s, e
}

func One(
	input string,
) (int, error) {
	g, s, e := newGrid(input)
	steps := paint(&g, e, s)
	return steps[s.row][s.col], nil
}

func paint(
	g *grid,
	zero coord,
	target coord,
) [41][161]int {
	max := len(g) * len(g[0])
	var output [41][161]int
	for i := range output {
		for j := range output[i] {
			output[i][j] = max
		}
	}
	output[zero.row][zero.col] = 0

	pending := make([]coord, 0, 8196)
	pending = append(pending, zero)

	var dest, s coord
	var dv int
	var msv, t uint8
	for len(pending) > 0 {
		dest = pending[0]
		if dest == target {
			break
		}

		dv = output[dest.row][dest.col] + 1

		msv = g[dest.row][dest.col] - 1
		if dest.row > 0 {
			s = coord{
				row: dest.row - 1,
				col: dest.col,
			}
			if g[s.row][dest.col] >= msv &&
				output[s.row][dest.col] > dv {

				output[s.row][s.col] = dv
				pending = append(pending, s)
			}
		}

		if t = dest.row + 1; t < numRows &&
			g[t][dest.col] >= msv &&
			output[t][dest.col] > dv {

			s = coord{
				row: t,
				col: dest.col,
			}
			output[s.row][s.col] = dv
			pending = append(pending, s)
		}

		if dest.col > 0 {
			s = coord{
				row: dest.row,
				col: dest.col - 1,
			}
			if g[s.row][s.col] >= msv &&
				output[s.row][s.col] > dv {

				output[s.row][s.col] = dv
				pending = append(pending, s)
			}
		}

		if t = dest.col + 1; t < numCols &&
			g[dest.row][t] >= msv &&
			output[dest.row][t] > dv {

			s = coord{
				row: dest.row,
				col: t,
			}
			output[s.row][s.col] = dv
			pending = append(pending, s)
		}

		pending = pending[1:]
	}

	return output
}
